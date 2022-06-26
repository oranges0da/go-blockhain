package block

import (
	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/proof"
	"github.com/oranges0da/goblockchain/tx"
)

// first find nonce using proof of work, then return hash of final block
func Hash(block *model.Block) (int, []byte) {
	pow := proof.New(block)
	nonce, hash := pow.Run()

	return nonce, hash[:]
}

func New(blockID int, prevHash []byte, tx *tx.Transaction) *model.Block {
	block := &model.Block{
		BlockID:     blockID,
		PrevHash:    prevHash,
		Transaction: tx,
	}

	return block
}

// like New() but only for genesis block, to param is the address the reward will be sent to
func Genesis(addr, msg string) *model.Block {
	coinbase := tx.NewCoinbase(addr, msg)

	block := &model.Block{
		PrevHash:    []byte("0"),
		BlockID:     0,
		Nonce:       0,
		Transaction: coinbase,
	}

	nonce, hash := Hash(block)
	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}
