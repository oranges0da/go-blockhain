package net

import "fmt"

func CmdToBytes(cmd string) []byte {
	var bytes [cmdLen]byte

	for i, c := range cmd {
		bytes[i] = byte(c)
	}

	return bytes[:]
}

func BytesToCmd(bytes []byte) string {
	var cmd []byte

	for _, b := range bytes {
		if b != 0x0 {
			cmd = append(cmd, b)
		}
	}

	return fmt.Sprintf("%v", cmd)
}

func ExtractCmd(request []byte) []byte {
	return request[:cmdLen]
}
