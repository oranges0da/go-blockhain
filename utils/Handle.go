package utils

import "log"

func Handle(err error) {
	if err != nil {
		log.Printf("Error handler: %v", err)
	}
}
