package model

import "github.com/dgraph-io/badger"

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
}
