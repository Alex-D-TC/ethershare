package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdsa"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
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

func generateRandomBytes(byteCount uint32) ([]byte, error) {
	arr := make([]byte, byteCount)
	_, err := rand.Read(arr)

	return arr, err
}

func encryptFileStreamAESCFB128(key []byte, ivSize uint32, chunkSize uint32, dst io.Writer, src io.Reader) error {
	if len(key) != 16 {
		return errors.New("The key is not of 16 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	chunk := make([]byte, chunkSize)

	for {
		readCount, err := src.Read(chunk)
		if err != nil {
			return err
		}

		// Generate a random IV
		iv, err := generateRandomBytes(ivSize)
		if err != nil {
			return err
		}

		encrypter := cipher.NewCFBEncrypter(block, iv)

		// If there is nothing more to encrypt, stop
		if readCount == 0 {
			break
		}

		// Encrypt the received chunk
		encryptedChunk := make([]byte, readCount)

		encrypter.XORKeyStream(encryptedChunk, chunk[0:readCount])

		// Append IV to chunk
		encryptedChunk = append(encryptedChunk, iv...)

		_, err = dst.Write(encryptedChunk)
		if err != nil {
			return err
		}

	}

	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Negotiate key, iv size and chunk size
	key := make([]byte, 16)

	// Get key as 128 bit byte array
	_, err := conn.Read(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send back the chunk and iv size
	chunkSize := uint32(64)
	ivSize := uint32(16)

	chunkSizeRawBuf := new(bytes.Buffer)
	ivSizeRawBuf := new(bytes.Buffer)

	// Serialize uint32 to [4]byte
	err = binary.Write(chunkSizeRawBuf, binary.LittleEndian, chunkSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the serialized bytes
	_, err = conn.Write(chunkSizeRawBuf.Bytes()[0:4])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Serialize uint32 to [4]byte
	err = binary.Write(ivSizeRawBuf, binary.LittleEndian, ivSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the serialized bytes
	_, err = conn.Write(ivSizeRawBuf.Bytes()[0:4])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the PDF to send
	// Hardcoded path for demonstration purposes
	pdfPath := "D:/projects/go/src/github.com/alex-d-tc/ethershare/CryptoFrontendProj/AVL-Tree-Rotations.pdf"
	file, err := os.Open(pdfPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//logMD5(pdfPath)

	// Stream the encrypted chunks
	encryptFileStreamAESCFB128(key, ivSize, chunkSize, conn, file)
}

func logMD5(path string) {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("MD5 Hash of pdf file: ", hash.Sum(nil))
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
