package proof

import (
	"bytes"
	"crypto/sha256"
	"fmt"
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

func (pow *PoW) PrepareData(nonce int64) []byte {
	blockData := [][]byte{
		utils.ToByte(int64(pow.Block.BlockID)),
		utils.ToByte(int64(pow.Block.Nonce)),
		utils.ToByte(nonce),
		pow.Block.Data,
	}

	data := bytes.Join(blockData, []byte{})

	return data
}

func (pow *PoW) Run() {
	var nonce int64 = 0
	var intHash big.Int

	for nonce < math.MaxInt {
		data := pow.PrepareData(nonce)
		hash := sha256.Sum256(data)

		intHash.SetBytes(hash[:])
		fmt.Printf("\r%x\n", hash)

		if intHash.Cmp(pow.Target) == -1 {
			log.Printf("Block found, hash: %x\n", hash)
			break
		} else {
			nonce++
		}
	}
}
