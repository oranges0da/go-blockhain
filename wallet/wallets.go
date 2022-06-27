package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"os"

	"github.com/oranges0da/goblockchain/handle"
)

const (
	walletFile = "tmp/wallet.data"
)

type Wallets struct {
	Wallets map[string]*Wallet // map of wallets, where key is address
}

func NewWallets() *Wallets {
	wallets := &Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	return wallets
}

func (ws *Wallets) Add(w *Wallet) error {
	ws.Wallets[w.Address] = w

	return nil
}

func (ws *Wallets) Get(addr string) *Wallet {
	return ws.Wallets[addr]
}

func (ws *Wallets) Save() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	enc := gob.NewEncoder(&content)
	err := enc.Encode(ws)
	handle.Handle(err, "Problem encoding/saving wallets.")

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	handle.Handle(err, "Problem saving wallets.")
}

func LoadWallets() *Wallets {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		handle.Handle(err, "Error loading wallets")
	}

	var wallets Wallets

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		handle.Handle(err, "Error loading wallets")
	}

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		handle.Handle(err, "Error loading wallets")
	}

	return &wallets
}
