package tx

import (
	"github.com/oranges0da/goblockchain/hash_utils"
	"github.com/oranges0da/goblockchain/model"
)

// checks that an address can unclock an input for spending
func InCanUnlock(in model.TxInput, addr string) bool {
	pubKeyHash := hash_utils.HashPubKey(in.PubKey)
	address := hash_utils.GetAddress(pubKeyHash)

	return address == addr
}
