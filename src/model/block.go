package model

type Block struct {
	BlockID     int
	Timestamp   int64
	Hash        []byte
	Nonce       int
	PrevHash    []byte
	Transaction *Transaction
}
