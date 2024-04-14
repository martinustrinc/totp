package util

import (
	"crypto/sha256"
	"encoding/hex"
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
