package model

import "github.com/dgraph-io/badger"

type BlockchainIter struct {
	CurrentHash []byte
	DB          *badger.DB
}
