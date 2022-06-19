package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/utils"
	"golang.org/x/crypto/ripemd160"
)

const (
	version = byte(0x00)
)

type Wallet struct {
	PrivKey ecdsa.PrivateKey
	PubKey  []byte
	Address string
	Balance int // balance in satoshis
}

func New() *Wallet {
	privKey, pubKey := NewKeyPair()
	wallet := Wallet{privKey, pubKey, "", 0}
	return &wallet
}

func (w Wallet) NewAddress() []byte {
	pubHash := HashPubKey(w.PubKey)

	versionedHash := append([]byte{version}, pubHash...)
	checkSum := CheckSum(versionedHash)

	fullHash := append(versionedHash, checkSum...)

	address := base58.Encode(fullHash)
}

// generate new private and public keys
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	utils.Handle(err, "Problem generating private key.")

	pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
	return *privKey, pubKey
}

func HashPubKey(pubKey []byte) []byte {
	hash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(hash[:])
	utils.Handle(err, "Problem hashing public key into address.")

	pubKeyHash := hasher.Sum(nil)

	return pubKeyHash
}

func CheckSum(hash []byte) []byte {
	first := sha256.Sum256(hash)
	second := sha256.Sum256(first[:])

	return second[:4]
}
