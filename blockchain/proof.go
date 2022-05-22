package blockchain

import (
	"fmt"
	"math/big"

	"github.com/oranges0da/go-blockchain/block"
)

const diff = 20

type Pow struct {
	Target *big.Int
	Block  *block.Block
}

func InitPow(block *block.Block) *Pow {
	target := big.NewInt(1)
	target.Lsh(target, 256-diff)

	pow := &Pow{
		Target: target,
		Block:  block,
	}

	return pow
}

func RunProofWork(chain *Blockchain) {
	pow := InitPow(chain.blocks[len(chain.blocks)-1])

	fmt.Printf("Mining the block containing \"%x\"\n", pow)
}

// func Validate() {}
