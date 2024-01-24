package util

import (
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"time"
)

func GeneratePassCode(secret string, timeNow time.Time) string {
	passcode, err := totp.GenerateCodeCustom(secret, timeNow, totp.ValidateOpts{
		Period:    120,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA256,
	})
	if err != nil {
		panic(err)
	}
	return passcode
}
