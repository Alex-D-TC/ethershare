package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/alex-d-tc/ethershare/eth/build-go"

	"github.com/alex-d-tc/ethershare/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alex-d-tc/ethershare/eth"
)

func deployTokenContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool) {

		auth, err := eth.PrepareTransactionAuth(client, key)
		if err != nil {
			fmt.Println(err)
			return err, true
		}

		addr, _, _, err := ethBind.DeployTokenImpl(auth, client, "FileShare", "FS", big.NewInt(1000000000))
		if err != nil {
			fmt.Println(err)
			return err, true
		}

		addrChan <- addr
		return nil, false
	})

	return <-addrChan
}

func deployFileshareContract(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient, tokenAddr common.Address) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool) {

		auth, err := eth.PrepareTransactionAuth(client, key)
		if err != nil {
			fmt.Println(err)
			return err, true
		}

		addr, _, _, err := ethBind.DeployFileShareContract(auth, client, tokenAddr)
		if err != nil {
			fmt.Println(err)
			return err, true
		}

		addrChan <- addr
		return nil, false
	})

	return <-addrChan
}

func main() {

	_, err := util.LoadKeys("./RES/eth.key")
	if err != nil {
		panic(err)
	}

	client, err := eth.GetThreadsafeClient("https://ropsten.infura.io/")
	if err != nil {
		panic(err)
	}
	defer func() {
		client.Close()
		client.Dispose()
	}()

	tokenAddr := common.HexToAddress("0xb0b8A1135Ceb0Eb672247Ac2AaCaE2169C9894c4")
	_ = common.HexToAddress("0x50b07d44523a7713BA4AeefA3867751f1cFD2657")

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
