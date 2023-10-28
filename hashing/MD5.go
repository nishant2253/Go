package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func md5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func main() {
	data := []byte("Hello, World!")
	result := md5Hash(data)
	fmt.Println(result)
}
