package util

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"time"
)

func TimeParse(timeLayout string, timeStr string) (timeRes time.Time, err error) {
	timeRes, err = time.Parse(timeLayout, timeStr)
	if err != nil {
		fmt.Println("Error -> Cannot Parse Time")
		return
	}
	return
}

func ByteArrayToString(dataByte []byte) string {
	byteSlice := dataByte[:]
	result := string(byteSlice)
	return result
}

func StringToBase64String(str string) (base64Str string) {
	// Convert string to byte slice
	strBytes := []byte(str)

	// Encode byte slice to base64 string
	base64Str = base64.StdEncoding.EncodeToString(strBytes)
	return
}

func StringToBase32String(str string) (base32Str string) {
	// Convert string to byte slice
	strBytes := []byte(str)

	// Encode byte slice to base64 string
	base32Str = base32.StdEncoding.EncodeToString(strBytes)
	return
}
