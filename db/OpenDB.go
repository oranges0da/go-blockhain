package db

import "github.com/xujiajun/nutsdb"

func OpenDB() (*nutsdb.DB, error) {
	opts := nutsdb.DefaultOptions
	opts.Dir = "tmp/blocks"

	db, err := nutsdb.Open(opts)

	if err != nil {
		return nil, err
	}
	return db, nil
}
