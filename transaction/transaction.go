package transaction

type Transaction struct {
	ID       []byte
	In       []TxInput
	Out      []TxOutput
	Locktime int64
}

type TxInput struct {
	Hash []byte // hash of transaction
	Sig  string
	Out  int // db index location of output of transaction
}

type TxOutput struct {
	Value  float64
	PubKey string // receiver's public key/address
}
