package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func sha1Hash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	data := []byte("Hello, World!")
	result := sha1Hash(data)
	fmt.Println(result)
}
