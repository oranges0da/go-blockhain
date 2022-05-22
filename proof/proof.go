package blockchain

import (
	"bytes"
	"crypto/sha256"
	"log"
	"math"
	"math/big"

	"github.com/oranges0da/go-blockchain/block"
	"github.com/oranges0da/go-blockchain/utils"
)

const diff = 12

type PoW struct {
	Target *big.Int
	Block  *block.Block
}

func New(block *block.Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, 256-diff)

	pow := &PoW{
		Target: target,
		Block:  block,
	}

	return pow
}

func (pow *PoW) PrepareData() []byte {
	blockData := [][]byte{
		utils.ToByte(pow.Block.BlockID),
		utils.ToByte(pow.Block.Nonce),
		pow.Block.Data,
	}

	data := bytes.Join(blockData, []byte{})

	return data
}

func (pow *PoW) Run() {
	nonce := 0
	var intHash big.Int

	for nonce < math.MaxInt {
		data := pow.PrepareData()
		hash := sha256.Sum256(data)

		log.Printf("\rMining current block, hash: %x\n", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			log.Printf("Block found, hash: %x\n", hash)
			break
		}
	}
}
