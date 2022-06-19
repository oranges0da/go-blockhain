package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"io/ioutil"
	"os"

	"github.com/oranges0da/goblockchain/utils"
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

func (ws *Wallets) AddWallet(w Wallet) error {
	ws.Wallets[w.Address] = &w

	return nil
}

func (ws *Wallets) GetWallet(addr string) *Wallet {
	return ws.Wallets[addr]
}

func (ws *Wallets) Save() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	enc := gob.NewEncoder(&content)
	err := enc.Encode(ws)
	utils.Handle(err, "Problem encoding/saving wallets.")

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	utils.Handle(err, "Problem saving wallets.")
}

func (ws *Wallets) Load() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	var wallets Wallets

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		return err
	}

	gob.Register(elliptic.P256())

	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		return err
	}

	ws.Wallets = wallets.Wallets

	return nil
}
