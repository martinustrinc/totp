package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
	"log"
	"os"
	"time"
	"totp-learn/constanta"
	"totp-learn/domain"
	"totp-learn/util"
)

func main() {
	var (
		err           error
		data          string
		secretTemp    string
		key           *otp.Key
		grochatBundle domain.RequestGrochatBundle
		args          = os.Args
		timeLayout    = constanta.TimeLayoutDefault
	)

	defer func() {
		if err != nil {
			log.Fatal("Error -> ", err.Error())
		}
	}()

	if len(args) < 1 {
		log.Fatal("Arguments Empty")
	}

	data = args[1]
	if err = json.Unmarshal([]byte(data), &grochatBundle); err != nil {
		fmt.Printf("Error -> Cannot Convert JSON\n")
		return
	}

	g := grochatBundle.DataDetail
	secretTemp = g.ClientID + g.HWID + g.TimestampStr + g.Type

	g.Timestamp, err = time.Parse(timeLayout, g.TimestampStr)
	if err != nil {
		return
	}

	if key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      "NEXSOFT-ND6",
		AccountName: "unknown@nexsoft.co.id",
		Algorithm:   otp.AlgorithmSHA256,
		Digits:      otp.DigitsSix,
		Secret:      []byte(util.CheckSumWithSha256([]byte(secretTemp))),
		SecretSize:  20,
	}); err != nil {
		fmt.Printf("Error -> Cannot Generate Key\n")
		return
	}

	otpPasscode := util.GeneratePassCode(key.Secret(), g.Timestamp)
	resultOtp := domain.ResponseOTP{OTP: otpPasscode}

	resp, err := json.Marshal(resultOtp)
	if err != nil {
		return
	}

	fmt.Println("Response -> ", string(resp))
	qrData, err := json.Marshal(g)
	if err != nil {
		return
	}

	//--- Prepare Data
	qrDataValue := flag.String("data", string(qrData), "Data For QR Code")
	qrOutputValue := flag.String("output", "qrcode.png", "File QR Code")

	flag.Parse()

	//--- Create QR code
	err = qrcode.WriteFile(*qrDataValue, qrcode.Medium, 256, *qrOutputValue)
	if err != nil {
		fmt.Printf("Error -> Cannot Generate QR Code")
		return
	}
}

func produceQR(grochatBundle domain.RequestGrochatBundle) (err error) {
	var (
		gro        = grochatBundle.DataDetail
		secretTemp = gro.ClientID + gro.HWID + gro.TimestampStr + gro.Type
		key        *otp.Key
	)

	gro.Timestamp, err = timeParse(constanta.TimeLayoutDefault, gro.TimestampStr)
	if err != nil {
		return
	}

	key, err = generateTotp(secretTemp)
	if err != nil {
		return
	}

	fmt.Println(key)
	return
}

func generateTotp(secret string) (key *otp.Key, err error) {
	key, err = totp.Generate(totp.GenerateOpts{
		Issuer:      "NEXSOFT-ND6",
		AccountName: "unknown@nexsoft.co.id",
		Algorithm:   otp.AlgorithmSHA256,
		Digits:      otp.DigitsSix,
		Secret:      []byte(util.CheckSumWithSha256([]byte(secret))),
		SecretSize:  20,
	})
	if err != nil {
		fmt.Println("Error -> Cannot Generate Key")
		return
	}
	return
}

func timeParse(timeLayout string, timeStr string) (timeRes time.Time, err error) {
	timeRes, err = time.Parse(timeLayout, timeStr)
	if err != nil {
		fmt.Println("Error -> Cannot Parse Time")
		return
	}
	return
}
