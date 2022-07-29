package block_utils

import "github.com/oranges0da/goblockchain/src/handle"

func GetBlockHashes() [][]byte {
	var block_hashes [][]byte

	blocks, err := GetBlocks()
	handle.Handle(err, "error getting blocks while getting block hashes")

	for _, b := range blocks {
		block_hashes = append(block_hashes, b.Hash)
	}

	return block_hashes
}
