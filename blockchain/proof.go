package blockchain

import (
	"math/big"

	"github.com/oranges0da/go-blockchain/block"
)

type ProofOfWork struct {
	Target *big.Int
	Block  *block.Block
}

func (chain *Blockchain) InitProofOfWork() *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, 256-64)

	pow := &ProofOfWork{
		Target: target,
		Block:  chain.blocks[len(chain.blocks)-1],
	}

	return pow
}
