package tx

import "github.com/oranges0da/goblockchain/hash_utils"

type TxInput struct {
	ID     []byte // hash of transaction that is being spent/consumed
	Vout   int    // index of output in the previous transaction that is being spent (Vector Output)
	Sig    []byte // signature of input
	PubKey []byte // pubkey of sender, used to sign and verify signature
}

// checks that an address can unclock an input for spending
func (in *TxInput) InCanUnlock(addr string) bool {
	pubKeyHash := hash_utils.HashPubKey(in.PubKey)
	address := hash_utils.GetAddress(pubKeyHash)

	return address == addr
}
