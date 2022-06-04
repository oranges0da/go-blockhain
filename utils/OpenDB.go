package utils

import "github.com/dgraph-io/badger"

func OpenDB() (*badger.DB, error) {
	opts := badger.DefaultOptions("/tmp/blocks")
	opts.Dir = "/tmp/blocks"
	opts.ValueDir = "/tmp/blocks/MANIFEST"
	opts.Logger = nil

	if db, err := badger.Open(opts); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
