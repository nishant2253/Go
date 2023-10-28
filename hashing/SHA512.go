package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func sha512Hash(data []byte) string {
	h := sha512.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	data := []byte("Hello, World!")
	result := sha512Hash(data)
	fmt.Println(result)
}
