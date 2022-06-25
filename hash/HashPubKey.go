package hash

import "crypto/sha256"

func HashPubKey(pubKey []byte) []byte {
	firstHash := sha256.Sum256(pubKey)
	finalHash := sha256.Sum256(firstHash[:])

	return finalHash[:]
}
