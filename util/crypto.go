package util

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func Sign(key *ecdsa.PrivateKey, hash []byte) ([]byte, error) {
	return crypto.Sign(hash[:], key)
}

func Verify(pubKeyRaw []byte, hashData []byte, signature []byte) bool {
	return crypto.VerifySignature(pubKeyRaw, hashData, signature[:64])
}

func PubKeysEqual(k1 ecdsa.PublicKey, k2 ecdsa.PublicKey) bool {
	return k1.X.Cmp(k2.X) == 0 && k1.Y.Cmp(k2.Y) == 0
}

func LoadKeys(path string) (*ecdsa.PrivateKey, error) {
	return crypto.LoadECDSA(path)
}

func WriteKeys(path string, privKey *ecdsa.PrivateKey) error {
	return crypto.SaveECDSA(path, privKey)
}

func MarshalPubKey(key *ecdsa.PublicKey) []byte {
	return crypto.FromECDSAPub(key)
}

func UnmarshalPubKey(rawKey []byte) (*ecdsa.PublicKey, error) {
	return crypto.UnmarshalPubkey(rawKey)
}
