package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func CheckSumWithSha256(content []byte) string {
	result := sha256.Sum256(content)
	return hex.EncodeToString(result[:])
}

func EncodeSHA256(input string) string {
	dataHashing := sha256.New()
	dataHashing.Write([]byte(input))
	return hex.EncodeToString(dataHashing.Sum(nil))
}

func DecodeSHA256(dataEncrypt string) (dataChipper []byte, err error) {
	dataChipper, err = hex.DecodeString(dataEncrypt)

	return
}

func EncryptData(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("error encrypt : ", err)
		return nil, err
	}

	// The IV needs to be unique, but not secure. Therefore, it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	//fmt.Println("Cipher text : ", ByteArrayToString(ciphertext))
	return ciphertext, nil
}

func DecryptData(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	//fmt.Println("Plain text : ", data)

	return data, nil
}
