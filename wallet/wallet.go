package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/utils"
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
	wallet := &Wallet{privKey, pubKey, "", 0}
	wallet.Address = wallet.NewAddress()

	return wallet
}

func ValidateAddress(addr string) bool {
	// decode and validate address by checking checksum
	decoded, err := base58.Decode(addr)
	utils.Handle(err, "Error decoding addres whilst validating.")

	// get checksum (last 4 bytes) from public key hash, but not the version (first byte)
	checkSum := decoded[len(decoded)-4:]
	pubKeyHash := decoded[1 : len(decoded)-4]

	testHash := sha256.Sum256(pubKeyHash)
	testCheckSum := testHash[:len(testHash)-4]

	return bytes.Equal(checkSum, testCheckSum)
}

// get address in base58 format from public key
func (w *Wallet) NewAddress() string {
	pubHash := HashPubKey(w.PubKey)

	versionedHash := append([]byte{version}, pubHash...)
	checkSum := CheckSum(versionedHash)

	fullHash := append(versionedHash, checkSum...)

	address := base58.Encode(fullHash)

	return address
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
	firstHash := sha256.Sum256(pubKey)
	finalHash := sha256.Sum256(firstHash[:])

	return finalHash[:]
}

// get 4-byte long checksum from pubHash
func CheckSum(hash []byte) []byte {
	first := sha256.Sum256(hash)
	second := sha256.Sum256(first[:])

	return second[:4]
}
