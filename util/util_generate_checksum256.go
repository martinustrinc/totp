package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func CheckSumWithSha256(content []byte) string {
	result := sha256.Sum256(content)
	return hex.EncodeToString(result[:])
}
