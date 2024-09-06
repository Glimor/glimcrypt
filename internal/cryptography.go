package internal

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"strings"
)

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func PKCS7UnPadding(data []byte) []byte {
	length := len(data)
	if length == 0 {
		fmt.Println("Invalid 8 keyword provided or file is not encrypted with GlimCrypt.")
		os.Exit(1)
	}
	unpadding := int(data[length-1])
	if unpadding > length || unpadding == 0 {
		fmt.Println("Invalid 8 keyword provided or file is not encrypted with GlimCrypt.")
		os.Exit(1)
	}
	return data[:(length - unpadding)]
}

func EncryptFile(filename string, key []byte) error {
	plaintext, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create cipher: %w", err)
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return fmt.Errorf("could not generate IV: %w", err)
	}

	plaintext = PKCS7Padding(plaintext, aes.BlockSize)

	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCBCEncrypter(block, iv)
	stream.CryptBlocks(ciphertext, plaintext)

	finalCiphertext := append(iv, ciphertext...)

	encryptedFile := filename + ".enc"
	return os.WriteFile(encryptedFile, finalCiphertext, 0644)
}

func DecryptFile(filename string, key []byte) error {
	ciphertext, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read encrypted file: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create cipher: %w", err)
	}

	if len(ciphertext) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCBCDecrypter(block, iv)
	stream.CryptBlocks(plaintext, ciphertext)

	plaintext = PKCS7UnPadding(plaintext)

	decryptedFile := strings.TrimSuffix(filename, ".enc") + ".dec"
	return os.WriteFile(decryptedFile, plaintext, 0644)
}
