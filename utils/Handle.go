package utils

import "log"

func Handle(err error) {
	if err != nil {
		log.Fatalf("Error handler: %v", err)
	}
}
