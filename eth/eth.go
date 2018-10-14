package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/alex-d-tc/bchain-routing/concurrent"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThreadsafeClient struct {
	sync.RWMutex

	ethClient            *ethclient.Client
	queue                *concurrent.TransactionQueue
	transactionWaitBlock chan bool
}

func MakeThreadsafeClient(client *ethclient.Client) *ThreadsafeClient {
	result := ThreadsafeClient{
		ethClient:            client,
		queue:                concurrent.MakeTransactionQueue(),
		transactionWaitBlock: make(chan bool, 1),
	}

	// Init chan with something, to not lock the first transaction
	result.transactionWaitBlock <- true

	return &result
}

func (client *ThreadsafeClient) SubmitTransaction(tran func(*ethclient.Client) (error, bool, *types.Transaction)) error {
	return client.queue.Submit(func() (error, bool) {

		// Wait for pending transaction to be mined
		<-client.transactionWaitBlock

		client.Lock()

		err, retry, tran := tran(client.ethClient)

		if err == nil {
			// Short enough deadline that no context leak should happen
			// No need for the cancel func for now
			ctx, _ := context.WithTimeout(context.Background(), 5*60*time.Second)

			fmt.Println(fmt.Sprintf("Waiting for transaction %s to be mined", tran.Hash().Hex()))

			bind.WaitMined(ctx, client.ethClient, tran)

			fmt.Println("Transaction mined or timer expired")
		}

		client.Unlock()

		// Notify that the next transaction can begin
		client.transactionWaitBlock <- true
		return err, retry
	})
}

func (client *ThreadsafeClient) SubmitReadTransaction(tran func(*ethclient.Client) (error, bool)) error {
	return client.queue.Submit(func() (error, bool) {

		client.RLock()

		err, cont := tran(client.ethClient)

		client.RUnlock()
		return err, cont
	})
}

func (client *ThreadsafeClient) Close() {
	client.ethClient.Close()
}

func (client *ThreadsafeClient) Dispose() {
	client.queue.Dispose()
}

func GetThreadsafeClient(rawUrl string) (*ThreadsafeClient, error) {
	client, err := GetClient(rawUrl)
	if err != nil {
		return nil, err
	}

	return MakeThreadsafeClient(client), nil
}

func GetClient(rawUrl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawUrl)
}

func PrepareTransactionAuth(client *ethclient.Client, key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(key)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	//auth.GasLimit = 500000

	return auth, nil
}

func EventWatcher(ctx context.Context, logger *log.Logger, client *ThreadsafeClient, filterProcessor func(*bind.FilterOpts)) {

	done := false
	lastBlock := uint64(0)

	for {

		select {
		case <-ctx.Done():
			done = true
			break
		default:
			break
		}

		if done {
			break
		}

		time.Sleep(5 * time.Second)

		client.RLock()

		header, err := client.ethClient.BlockByNumber(ctx, nil)
		if err != nil {
			logger.Println(err)
			// Retry on error
			continue
		}

		client.RUnlock()

		block := uint64(header.Number().Int64())

		if lastBlock == block {
			//fmt.Println("Block already processed. Sleeping")
			continue
		}

		//logger.Println(fmt.Sprintf("Checking events from block %d to block %d", lastBlock+1, block))

		opts := &bind.FilterOpts{
			Context: ctx,
			Start:   lastBlock + 1,
			End:     &block,
		}

		filterProcessor(opts)
		lastBlock = block
	}
}
