package model

import "github.com/oranges0da/goblockchain/transaction"

type Block struct {
	BlockID     int
	PrevHash    []byte
	Nonce       int
	Transaction *transaction.Transaction
	Hash        []byte
}
