package block

import (
	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/proof"
	"github.com/oranges0da/goblockchain/transaction"
)

// first find nonce using proof of work, then return hash of final block
func HashBlock(b *model.Block) (int, []byte) {
	pow := proof.New(b)
	nonce, hash := pow.Run()

	return nonce, hash[:]
}

func New(BlockId int, prevHash []byte, txs []*transaction.Transaction) *model.Block {
	block := &model.Block{
		BlockID:      BlockId,
		PrevHash:     []byte{},
		Transactions: txs,
	}

	nonce, hash := HashBlock(block)
	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}

// like New() but only for genesis block, to param is the address the reward will be sent to
func Genesis(to string) *model.Block {
	var transactions []*transaction.Transaction

	coinbase := transaction.NewCoinbase(to, "example sig")
	transactions = append(transactions, coinbase)

	block := &model.Block{
		PrevHash:     []byte("0"),
		BlockID:      0,
		Nonce:        0,
		Transactions: transactions,
	}

	nonce, hash := HashBlock(block)
	block.Nonce = nonce
	block.Hash = hash[:]

	return block
}
