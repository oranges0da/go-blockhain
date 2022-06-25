package model

import "github.com/oranges0da/goblockchain/tx"

type Block struct {
	BlockID     int
	Timestamp   int64
	Hash        []byte
	Nonce       int
	PrevHash    []byte
	Transaction *tx.Transaction
}
