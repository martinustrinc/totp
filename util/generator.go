package util

import (
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"log"
	"time"
)

// GenerateOtpCustom Generate TOTP custom
func GenerateOtpCustom(secret string, timeValue time.Time) (result string, err error) {
	//fmt.Println("Time Generate OTP : ", timeValue)
	otpCode, errs := totp.GenerateCodeCustom(secret, timeValue, totp.ValidateOpts{
		Period:    120,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})
	if errs != nil {
		//panic(err)
		//fmt.Printf("Error -> Cannot Generate OTP with error \n")
		fmt.Println("Error -> Cannot Generate OTP with error : ", errs)
		err = errs
		return
	}

	return otpCode, err
}

// GenerateOtpBasic Generate TOTP basic
func GenerateOtpBasic(secret string, timeValue time.Time) (result string, err error) {
	otpCode, errs := totp.GenerateCode(secret, timeValue)
	if errs != nil {
		return result, errs
	}
	return otpCode, err
}

// ValidateOtpCustom Validate TOTP custom
func ValidateOtpCustom(secret string, otpCode string, timeValue time.Time) (valid bool, err error) {
	valid, err = totp.ValidateCustom(otpCode, secret, timeValue, totp.ValidateOpts{
		Period:    120,
		Skew:      1,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})
	if err != nil {
		log.Println("Error in ValidateOtpCustom : ", err)
	}
	return
}

// ValidateOtpBasic Validate TOTP basic
func ValidateOtpBasic(secret string, otpCode string) bool {
	return totp.Validate(otpCode, secret)
}

func GenerateOtpSecretKey(dataInput []byte) (result string, err error) {
	// Generate a new TOTP key
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "NexSOFT-TOTP",
		AccountName: "nextrac@nexcloud.co.id",
		Algorithm:   otp.AlgorithmSHA256,
		Digits:      otp.DigitsSix,
		Secret:      []byte(CheckSumWithSha256(dataInput)),
		SecretSize:  20,
	})
	if err != nil {
		fmt.Printf("Error -> Cannot Generate Key with error \n")
		return result, err
	}
	secretKey := key.Secret()

	return secretKey, err
}
