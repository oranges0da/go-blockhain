package proof

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"

	"github.com/oranges0da/goblockchain/model"
	"github.com/oranges0da/goblockchain/utils"
)

const diff = 12

type PoW struct {
	Target *big.Int // hash target that should be reached with nonce
	Block  *model.Block
}

func New(block *model.Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, 256-diff) // make new hash target with difficulty

	pow := &PoW{
		Target: target,
		Block:  block,
	}

	return pow
}

func (pow *PoW) PrepareData(nonce int) []byte {
	blockData := [][]byte{
		utils.ToByte(int(pow.Block.BlockID)),
		utils.ToByte(int(pow.Block.Nonce)),
		utils.ToByte(nonce),
		utils.ToByte(pow.Block.Transactions),
	}

	data := bytes.Join(blockData, []byte{})

	return data
}

func (pow *PoW) Run() (int, [32]byte) {
	var hash [32]byte
	var nonce int = 0
	var intHash big.Int

	for nonce < math.MaxInt { // for every number in range run hash to find hash target
		data := pow.PrepareData(nonce)
		testHash := sha256.Sum256(data)

		intHash.SetBytes(testHash[:])
		fmt.Printf("\r Trying hash: %x", testHash)
		fmt.Printf("/r For block: %v", pow.Block)

		if intHash.Cmp(pow.Target) == -1 {
			fmt.Printf("\n")
			fmt.Printf("Block found, hash: %x\n", testHash)
			hash = testHash
			break
		} else {
			nonce++
		}
	}

	return nonce, hash
}
