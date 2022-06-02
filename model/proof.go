package model

import "math/big"

type PoW struct {
	Target *big.Int // hash target that should be reached with nonce
	Block  *Block
}
