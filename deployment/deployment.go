package deployment

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/alex-d-tc/ethershare/eth"
	"github.com/alex-d-tc/ethershare/eth/ethBind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployTokenContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool, *types.Transaction) {

		auth, err := eth.PrepareTransactionAuth(client, key)
		if err != nil {
			fmt.Println(err)
			return err, true, nil
		}

		addr, tran, _, err := ethBind.DeployTokenImpl(auth, client, "FileShare", "FS", big.NewInt(1000000000))
		if err != nil {
			fmt.Println(err)
			return err, true, nil
		}

		addrChan <- addr
		return nil, false, tran
	})

	return <-addrChan
}

func DeployFileshareContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient, tokenAddr common.Address) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool, *types.Transaction) {

		auth, err := eth.PrepareTransactionAuth(client, key)
		if err != nil {
			fmt.Println(err)
			return err, true, nil
		}

		addr, tran, _, err := ethBind.DeployFileShareContract(auth, client, tokenAddr)
		if err != nil {
			fmt.Println(err)
			return err, true, nil
		}

		addrChan <- addr
		return nil, false, tran
	})

	return <-addrChan
}
