package util

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"strings"
)

func GetPrivateIP() (string, error) {

	con, err := net.Dial("udp", "8.8.8.8:80")
	defer con.Close()

	if err != nil {
		return "", err
	}

	addrSplit := strings.Split(con.LocalAddr().String(), ":")

	return addrSplit[0], nil
}

func GetPublicIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	resultBytes := bytes.NewBuffer([]byte{})
	_, err = io.Copy(resultBytes, resp.Body)

	return string(resultBytes.Bytes()), err
}
