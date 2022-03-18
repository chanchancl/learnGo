package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {

	plaintext := []byte("abcdefg")
	key, _ := hex.DecodeString("6368616e676520746869732070617373")

	block, _ := aes.NewCipher(key)

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	io.ReadFull(rand.Reader, iv)
	fmt.Println(ciphertext, iv)

	stream := cipher.NewCFBEncrypter(block, iv)

	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	fmt.Println(ciphertext, iv)
	fmt.Println(ciphertext[aes.BlockSize:])
}
