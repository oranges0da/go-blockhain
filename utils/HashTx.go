package utils

import (
	"crypto/sha256"

	"github.com/oranges0da/goblockchain/transaction"
)

func HashTransaction(tx *transaction.Transaction) [32]byte {
	byte_data := ToByte(tx)

	hash := sha256.Sum256(byte_data)

	return hash
}
