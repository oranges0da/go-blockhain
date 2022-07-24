package handle

import "log"

func Handle(err error, msg string) {
	if err != nil {
		log.Printf("Error caught whilst handling: %v", err)
		log.Printf("Error msg: %s", msg)
	}
}
