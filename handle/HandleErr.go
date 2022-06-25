package handle

import "log"

func Handle(err error, msg string) {
	log.Printf("Error handled: %v", err)
	log.Printf("Error msg: %s", msg)
}
