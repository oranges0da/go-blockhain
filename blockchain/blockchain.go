package blockchain

type Blockchain struct {
	LastHash []byte
	blocks   []*Block
}
