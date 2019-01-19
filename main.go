package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/alex-d-tc/ethershare/util"
)

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

/*
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
*/

func encryptFileStreamAESCFB256(key []byte, chunkSize uint32, dst io.Writer, src io.Reader) error {
	if len(key) != 32 {
		return errors.New("The key is not of 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := key[0:16]
	encrypter := cipher.NewCFBEncrypter(block, iv)

	chunk := make([]byte, chunkSize)

	for {
		readCount, err := src.Read(chunk)
		if err != nil {
			return err
		}

		// If there is nothing more to encrypt, stop
		if readCount == 0 {
			break
		}

		// If we read less than chunkSize, zero the rest
		for i := uint32(readCount); i < chunkSize; i++ {
			chunk[i] = 0
		}

		/*
			chunk, err = pkcs7Pad(chunk, block.BlockSize())
			if err != nil {
				return err
			}
		*/

		encryptedChunk := make([]byte, len(chunk))

		encrypter.XORKeyStream(encryptedChunk, chunk)
		_, err = dst.Write(encryptedChunk)
		if err != nil {
			return err
		}

		//fmt.Println(encryptedChunk)
	}

	return nil
}

func handleConnection(c net.Conn) {

	helloWorld := "Hello World!"
	key := make([]byte, 32)

	encryptFileStreamAESCFB256(key, 1024, c, bytes.NewReader([]byte(helloWorld)))
	c.Close()
}

func requestAndDecrypt() {

	key := make([]byte, 32)
	iv := key[0:16]

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp4", "localhost:8080")
	if err != nil {
		panic(err)
	}

	src := make([]byte, 1024)

	_, err = conn.Read(src)
	if err != nil {
		panic(err)
	}

	conn.Close()

	result := make([]byte, 1024)

	decrypter := cipher.NewCFBDecrypter(block, iv)
	decrypter.XORKeyStream(result, src)

	fmt.Println(result)
}

// pkcs7Pad right-pads the given byte slice with 1 to n bytes, where
// n is the block size. The size of the result is x times n, where x
// is at least 1.
func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, errors.New("Invalid block size")
	}
	if b == nil || len(b) == 0 {
		return nil, errors.New("Invalid PKCS7 data")
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

func main() {
	PORT := ":8080"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
