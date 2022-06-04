package utils

import (
	"log"
)

func Handle(err error, pkg string) {
	if err != nil {
		log.Printf("Package: %v \n Handler:%v ", pkg, err)
	}
}
