package blockchainDeployment

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/alex-d-tc/ethershare/blockchainEth"
	"github.com/alex-d-tc/ethershare/blockchainEth/ethBind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DeployTokenContract(key *ecdsa.PrivateKey, client *blockchainEth.ThreadsafeClient) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool, *types.Transaction) {

		auth, err := blockchainEth.PrepareTransactionAuth(client, key)
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

func DeployFileshareContract(key *ecdsa.PrivateKey, client *blockchainEth.ThreadsafeClient, tokenAddr common.Address) common.Address {

	addrChan := make(chan common.Address)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool, *types.Transaction) {

		auth, err := blockchainEth.PrepareTransactionAuth(client, key)
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
