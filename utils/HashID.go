package utils

import (
	"crypto/sha256"

	"github.com/oranges0da/goblockchain/transaction"
)

func HashID(tx *transaction.Transaction) [32]byte {
	byte_tx := ToByte(tx)

	hash := sha256.Sum256(byte_tx)

	return hash
}
