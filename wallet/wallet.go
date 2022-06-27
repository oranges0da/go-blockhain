package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/hash_utils"
)

type Wallet struct {
	PrivKey    ecdsa.PrivateKey
	PubKey     []byte
	PubKeyHash []byte
	Address    string
	Balance    int // balance in satoshis
}

func New() *Wallet {
	privKey, pubKey := NewKeyPair()
	pubKeyHash := hash_utils.HashPubKey(pubKey)
	address := hash_utils.GetAddress(pubKeyHash)

	wallet := &Wallet{
		PrivKey:    privKey,
		PubKey:     pubKey,
		PubKeyHash: pubKeyHash,
		Address:    address,
	}

	return wallet
}

func ValidateAddress(addr string) bool {
	// decode and validate address by checking checksum
	decoded, err := base58.Decode(addr)
	if err != nil {
		panic(err)
	}

	// get checksum (last 4 bytes) from public key hash, but not the version (first byte)
	checkSum := decoded[len(decoded)-4:]
	pubKeyHash := decoded[1 : len(decoded)-4]

	testHash := sha256.Sum256(pubKeyHash)
	testCheckSum := testHash[:len(testHash)-4]

	return bytes.Equal(checkSum, testCheckSum)
}

// generate new private and public keys
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		panic(err)
	}

	pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	return *privKey, pubKey
}
