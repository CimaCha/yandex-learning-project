package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key, err := generateRandomBytes(6)
	if err != nil {
		panic(err)
	}
	fmt.Println(key)
}

func generateRandomBytes(n int) (string, error) {
	// определяем слайс байт нужной длины
	b := make([]byte, n)
	_, err := rand.Read(b) // записываем байты в слайс b
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}
