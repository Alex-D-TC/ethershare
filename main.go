package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/alex-d-tc/ethershare/eth/build-go"

	"github.com/alex-d-tc/ethershare/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alex-d-tc/ethershare/eth"
)

func deployTokenContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient) common.Address {

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

func deployFileshareContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient, tokenAddr common.Address) common.Address {

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

type File struct {
	Hash  string
	Price *big.Int
}

func listFiles(client *eth.ThreadsafeClient, fileshareContractAddr common.Address, sharer common.Address) []File {
	results := []File{}

	done := make(chan bool)

	client.SubmitReadTransaction(func(client *ethclient.Client) (err error, retry bool) {
		err = nil
		retry = false

		fileshareContract, err := ethBind.NewFileShareContract(fileshareContractAddr, client)
		if err != nil {
			done <- true
			return
		}

		count, err := fileshareContract.SharesCount(nil, sharer)
		if err != nil {
			done <- true
			return
		}

		for idx := big.NewInt(0); count.Cmp(idx) > 0; idx = idx.Add(idx, big.NewInt(1)) {
			result, err := fileshareContract.Shares(nil, sharer, idx)

			// Just stop prematurely for now
			// TODO: Maybe just skip it in future versions?
			if err != nil {
				done <- true
				return err, false
			}

			results = append(results, File{
				Hash:  result.FileHash,
				Price: result.Price,
			})
		}

		done <- true
		return
	})

	// Wait for the transaction to be done
	<-done

	return results
}

func main() {

	_, err := util.LoadKeys("./RES/eth.key")
	if err != nil {
		panic(err)
	}

	client, err := eth.GetThreadsafeClient("https://rinkeby.infura.io/")
	if err != nil {
		panic(err)
	}
	defer func() {
		client.Close()
		client.Dispose()
	}()

	tokenAddr := common.HexToAddress("0xbD20ED28Cb9C0bbeb19Ff9709489B591f97250D9")
	_ = common.HexToAddress("0xd5967339f0A2Cb2E151A42AFF73e7cc7B7d631B1")

	done := make(chan bool)

	client.SubmitReadTransaction(func(client *ethclient.Client) (error, bool) {
		token, err := ethBind.NewTokenImpl(tokenAddr, client)
		if err != nil {
			done <- true
			return err, false
		}

		total, err := token.MaxSupply(nil)
		if err != nil {
			done <- true
			return err, false
		}

		fmt.Println(total.Text(10))

		owner, err := token.Owner(nil)
		if err != nil {
			done <- true
			return err, false
		}

		fmt.Println(owner.Hex())

		done <- true
		return nil, false
	})

	<-done
}
