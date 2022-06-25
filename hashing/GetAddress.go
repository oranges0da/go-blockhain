package hashing

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
)

const (
	version = byte(0x00)
)

// returns address encoded in base58 with checksum from provided pubKey
func GetAddress(pubKey []byte) string {
	var addr []byte

	// hash public key using simple func
	pubKeyHash := HashPubKey(pubKey)

	checkSumHash := sha256.Sum256(pubKeyHash)
	// get actual checksum from previous hash
	checkSum := checkSumHash[len(checkSumHash)-4:]

	// append version byte, pubKeyHash, and checkSum byte array
	addr = append(addr, version)
	addr = append(addr, pubKeyHash...)
	addr = append(addr, checkSum...)

	// encode full address in base58 for readability and useability
	address := base58.Encode(addr)

	return address
}
