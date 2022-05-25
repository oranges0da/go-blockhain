package transaction

type Transaction struct {
	In       []TxInput
	Out      []TxOutput
	Locktime int64
	Hash     []byte
}

type TxInput struct {
	ID  []byte // id of transaction pointing to correlated output in badger db
	Sig string
	Out int // db index location of output of transaction
}

type TxOutput struct {
	Value  float64
	PubKey string // receiver's public key/address
}

func NewCoinbase(to string, sig string) *Transaction {
	tx := &Transaction{
		ID:       []byte{},
		In:       []TxInput{},
		Out:      []TxOutput{},
		Locktime: 0,
	}

	return tx
}
