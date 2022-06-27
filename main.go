package main

import (
	"fmt"

	"github.com/oranges0da/goblockchain/blockchain"
	"github.com/oranges0da/goblockchain/handle"
	"github.com/oranges0da/goblockchain/wallet"
)

func main() {
	BlockChain, err := blockchain.New("randomaddress")
	handle.Handle(err, "Error creating blockchain in main.")
	fmt.Printf("Blockchain: %v", BlockChain)

	w := wallet.New()
	wallets := wallet.NewWallets()

	wallets.Add(w)
	wallets.Save()

	wal := wallets.Get("1SHf8PShy4ryFypVCrKt8QRb5JyJcSXeXPCnHKm1sTTWYy5eVS")

	fmt.Printf("wallet: %v", wal)
}
