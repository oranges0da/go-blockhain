package model

import "github.com/oranges0da/goblockchain/transaction"

type Block struct {
	BlockID      int
	PrevHash     [32]byte
	Nonce        int
	Transactions []*transaction.Transaction
	Hash         [32]byte
}
