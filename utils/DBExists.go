package utils

import "os"

func DBExists() bool {
	if _, err := os.Stat("/tmp/blocks"); os.IsNotExist(err) {
		return false
	}

	return true
}
