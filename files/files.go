package files

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/alex-d-tc/ethershare/eth"
	"github.com/alex-d-tc/ethershare/eth/ethBind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type File struct {
	Hash        string
	Description string
	Price       *big.Int
}

func PublishFile(key *ecdsa.PrivateKey, client *eth.ThreadsafeClient, fileshareContractAddr common.Address, file File) error {

	err := client.SubmitTransaction(func(client *ethclient.Client) (err error, retry bool, tran *types.Transaction) {

		err = nil
		retry = true

		fileshareContract, err := ethBind.NewFileShareContract(fileshareContractAddr, client)
		if err != nil {
			retry = false
			return
		}

		auth, err := eth.PrepareTransactionAuth(client, key)
		if err != nil {
			return
		}

		tran, err = fileshareContract.PublishFile(auth, file.Hash, file.Price, file.Description)
		if err != nil {
			return
		}

		return
	})

	if err != nil {
		return err
	}

	return nil
}

func ListFiles(client *eth.ThreadsafeClient, fileshareContractAddr common.Address, sharer common.Address) ([]File, error) {
	results := []File{}

	err, done := client.SubmitReadTransactionWait(func(client *ethclient.Client) (err error, retry bool) {
		err = nil
		retry = false

		fileshareContract, err := ethBind.NewFileShareContract(fileshareContractAddr, client)
		if err != nil {
			return
		}

		count, err := fileshareContract.SharesCount(nil, sharer)
		if err != nil {
			return
		}

		for idx := big.NewInt(0); count.Cmp(idx) > 0; idx = idx.Add(idx, big.NewInt(1)) {
			result, err := fileshareContract.Shares(nil, sharer, idx)

			// Just stop prematurely for now
			// TODO: Maybe just skip it in future versions?
			if err != nil {
				return err, false
			}

			results = append(results, File{
				Hash:  result.FileHash,
				Price: result.Price,
			})
		}

		return
	})

	if err != nil {
		return nil, err
	}

	// Wait for the transaction to be done
	<-done

	return results, nil
}
