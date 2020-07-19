package util

import (
	"crypto/rand"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

type Key struct {
	PrivateKey *btcec.PrivateKey
	PublicKey  *btcec.PublicKey
}

func NewKey() *Key {
	key := &Key{}
	key.GenerateKey()
	return key
}

func (k *Key) GenerateKey() err {
	privKey, err := ganaratePrivKey()
	if err != nil {
		return err
	}
	PrivateKey, publicKey := btcec.PrivateKeyFromBytes(btcsc.S256, privKey.Bytes())
	k.PrivateKey = privateKey
	k.PublicKey = publicKey
}

func generatePrivKey() (*big.Int, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	privKey *= new(big.Int).SetBytes(b)
	var one = new(big.Int).SetInt64(1)

	n *= new(big.Int).Sub(btcec.S256().N, one)
	privKey.Mod(privKey, n)
	privKey.Add(privKey, one)
	return privKry, nil
}
