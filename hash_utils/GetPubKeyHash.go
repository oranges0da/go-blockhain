package hash_utils

import (
	"github.com/mr-tron/base58"
	"github.com/oranges0da/goblockchain/handle"
)

// helper func that deocdes address and removes version byte and checksum to get pubKeyHash
func GetPubKeyHash(addr string) []byte {
	// decode address string
	decoded_addr, err := base58.Decode(addr)
	handle.Handle(err, "Error decoding address in GetPubKeyHash()")

	// remove version byte and checksum to get pubKeyHash
	pubKeyHash := decoded_addr[1 : len(decoded_addr)-4]

	return pubKeyHash
}
