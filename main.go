package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/alex-d-tc/ethershare/blockchainEth"
	"github.com/alex-d-tc/ethershare/blockchainEth/ethBind"
	"github.com/alex-d-tc/ethershare/blockchainFilesManagement"
	"github.com/alex-d-tc/ethershare/util"
)

func serverExample(client *blockchainEth.ThreadsafeClient, fileshareAddr common.Address) {

	r := gin.Default()

	r.GET("/echo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "echo!",
		})
	})

	r.GET("/:sharer/files", func(c *gin.Context) {
		sharerHex := c.Param("sharer")
		sharerAddress := common.HexToAddress(sharerHex)

		files, err := blockchainFilesManagement.ListFiles(client, fileshareAddr, sharerAddress)

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

func prepareConfig() (key *ecdsa.PrivateKey, clientURL string, tokenAddr common.Address, fileshareAddr common.Address, err error) {

	key, err = util.LoadKeys("./RES/eth.key")
	if err != nil {
		return
	}

	clientURL = "https://rinkeby.infura.io/"

	// Token contract Address
	tokenAddr = common.HexToAddress("0xbD20ED28Cb9C0bbeb19Ff9709489B591f97250D9")

	// Fileshare contract Address
	fileshareAddr = common.HexToAddress("0xd5967339f0A2Cb2E151A42AFF73e7cc7B7d631B1")

	return
}

func main() {

	key, clientURL, tokenAddr, _, err := prepareConfig()
	if err != nil {
		panic(err)
	}

	client, err := blockchainEth.GetThreadsafeClient(clientURL)
	if err != nil {
		panic(err)
	}
	defer func() {
		client.Close()
		client.Dispose()
	}()

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
