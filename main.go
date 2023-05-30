package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
)

func decrypt(key, iv, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	blockMode.CryptBlocks(plaintext, data)

	return plaintext, nil
}

func main() {
	key, err := hex.DecodeString("0602000000a400005253413100040000")
	if err != nil {
		log.Fatal(err)
	}

	iv, err := hex.DecodeString("0100010067244F436E6762F25EA8D704")
	if err != nil {
		log.Fatal(err)
	}

	hexStrCipher := "889df1f5802774a5d245be78b17e56a01f16128664883e73b9025e7b782e0f7eb061f1697ba9aa4641f1cc27519773e74e58e5f208abb64a8ee1b0f6e4770278"
	ciphertext, err := hex.DecodeString(hexStrCipher)
	if err != nil {
		log.Fatal(err)
	}

	plaintext, err := decrypt(key, iv, ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", plaintext)

	password := string(plaintext)
	fmt.Println(password)
}
