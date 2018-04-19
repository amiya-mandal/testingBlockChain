package main

import "fmt"

func IntToHex(num int64) []byte {
	s := fmt.Sprintf("%0x", num)
	return []byte(s)
}
