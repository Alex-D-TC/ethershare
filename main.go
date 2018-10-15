package main

import (
	"fmt"

	"github.com/alex-d-tc/ethershare/eth/ethBind"
	"github.com/alex-d-tc/ethershare/files"
	"github.com/alex-d-tc/ethershare/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/alex-d-tc/ethershare/eth"
)

func serverExample(client *eth.ThreadsafeClient, fileshareAddr common.Address) {

	r := gin.Default()

	r.GET("/echo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "echo!",
		})
	})

	r.POST("/:sharer/files", func(c *gin.Context) {

	})

	r.GET("/:sharer/files", func(c *gin.Context) {
		sharerHex := c.Param("sharer")
		sharerAddress := common.HexToAddress(sharerHex)

<<<<<<< HEAD
type File struct {
	Hash  string
	Price *big.Int
}

func getFile(client *eth.ThreadsafeClient, fileshareContractAddr common.Address, sharer common.Address, fileID *big.Int) (File, error) {

	result := File{}

	err, done := client.SubmitReadTransactionWait(func(client *ethclient.Client) (err error, retry bool) {

		err = nil
		retry = false

		fileshareContract, err := ethBind.NewFileShareContract(fileshareContractAddr, client)
		if err != nil {
			return
		}

		fileStruct, err := fileshareContract.Shares(nil, sharer, fileID)
		if err != nil {
			return
		}

		result = File{
			Hash:  fileStruct.FileHash,
			Price: fileStruct.Price,
		}

		return
	})

	if err != nil {
		return File{}, err
	}

	<-done
	return result, nil
}

func listFiles(client *eth.ThreadsafeClient, fileshareContractAddr common.Address, sharer common.Address) ([]File, error) {
	results := []File{}

	err, done := client.SubmitReadTransactionWait(func(client *ethclient.Client) (err error, retry bool) {
		err = nil
		retry = false

		fileshareContract, err := ethBind.NewFileShareContract(fileshareContractAddr, client)
=======
		files, err := files.ListFiles(client, fileshareAddr, sharerAddress)
>>>>>>> 493b0f3cf789ddbf8f693d4ffbfc1be1b6fa0c0c
		if err != nil {
			c.Error(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"files": files,
			})
		}
	})

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func main() {

	key, err := util.LoadKeys("./RES/eth.key")
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

	// Token contract Address
	tokenAddr := common.HexToAddress("0xbD20ED28Cb9C0bbeb19Ff9709489B591f97250D9")

	// Fileshare contract Address
	_ = common.HexToAddress("0xd5967339f0A2Cb2E151A42AFF73e7cc7B7d631B1")

	err, done := client.SubmitReadTransactionWait(func(client *ethclient.Client) (error, bool) {
		token, err := ethBind.NewTokenImpl(tokenAddr, client)
		if err != nil {
			return err, false
		}

		total, err := token.MaxSupply(nil)
		if err != nil {
			return err, false
		}

		fmt.Println(total.Text(10))

		owner, err := token.Owner(nil)
		if err != nil {
			return err, false
		}

		fmt.Println(owner.Hex())
		fmt.Println(crypto.PubkeyToAddress(key.PublicKey).Hex())

		return nil, false
	})

	if err != nil {
		panic(err)
	}

	<-done
}
