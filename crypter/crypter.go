package crypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Crypter struct {
	Key string
}

func NewCrypter() *Crypter {
	key := os.Getenv("KEY")

	if key == "" {
		panic("Нет ключа в параметре окружения")
	}

	return &Crypter{
		Key: key,
	}
}

func (c *Crypter) Encrypt(data []byte) []byte {
	block, err := aes.NewCipher([]byte(c.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}

	return aesGCM.Seal(nonce, nonce, data, nil)
}

func (c *Crypter) Decrypt(data []byte) []byte {
	block, err := aes.NewCipher([]byte(c.Key))
	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	text, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}

	return text
}
