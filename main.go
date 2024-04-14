package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	dtoModel "totp-learn/model"
	"totp-learn/services"
)

func main() {
	var (
		err        error
		dataBundle dtoModel.RequestDataBundleOtp
		secretKey  string
		timeOtp    time.Time
		clientID   string
	)

	//fmt.Println("Encode Data Bundle : ", util.EncodeSHA256("47e8d0327c1a43cfa1f83195aa60db66_B-/1HRD8T2/CNCMC0092J00A7/ - Audentis Fortuna iuvat_2024-04-09 15:57:50.903_ND6-sysadminLv1"))

	// Initialize SecretKeyStore
	keyStore := services.NewSecretKeyStore()
	secretKey = "your_secret_key"

	// Set secret keys for users
	//keyStore.SetSecretKey("user1", "secretKey1")
	//keyStore.SetSecretKey("user2", "secretKey2")

	defer func() {
		if err != nil {
			log.Fatal("Error -> ", err.Error())
		}
	}()

	//read argument action when call generator
	if len(os.Args) < 1 {
		fmt.Println("Usage: generator <action> [args...]")
		return
	}

	action := os.Args[1]
	if len(action) < 1 {
		log.Fatal("Arguments Empty")
	}

	//data = args[1]
	//if err = json.Unmarshal([]byte(data), &grochatBundle); err != nil {
	//	fmt.Printf("Error -> Cannot Convert JSON\n")
	//	return
	//}

	switch action {
	//To Generate OTP call or run with (./generator generate <dataBundle>)
	case "generate":
		if len(os.Args) < 2 {
			fmt.Println("Usage: <dirPath> generate <dataBundle>")
			return
		}

		dataBundleEncrypt := os.Args[2]
		//fmt.Println("action Command 0 : ", os.Args[0])
		//fmt.Println("action Command 1 : ", os.Args[1])
		//fmt.Println("action Command 2 : ", os.Args[2])

		//Decode secretKey
		//dataChipper, err := util.DecodeSHA256(dataBundleEncrypt)
		//if err != nil {
		//	fmt.Println("Error Decode Data Bundle :", err.Error())
		//	return
		//}
		//Generate secretKey
		//secretKey, err = util.GenerateOtpSecretKey(dataChipper)

		//Store data bundle in struct
		//dataCheckSum := util.CheckSumWithSha256(dataChipper)
		//dataBundleString := util.ByteArrayToString(dataChipper)
		//fmt.Println("Data bundle string = ", dataBundleEncrypt)
		dataSplit := strings.Split(dataBundleEncrypt, "_")
		//dataSplit := strings.Split(dataBundleEncrypt, "_")
		dataRequest := dataBundle.DataDetail
		if len(dataSplit) < 2 {
			fmt.Println("Error Data Bundle Invalid!")
			return
		}
		if len(dataSplit) > 3 {
			dataRequest.ClientID = dataSplit[0]
			dataRequest.HwID = dataSplit[1]
			dataRequest.TimestampStr = dataSplit[2]
			dataRequest.Type = dataSplit[3]
		}
		//Store clientId into global var
		clientID = dataRequest.ClientID

		//Store secretKey
		keyStore.SetSecretKey(clientID, secretKey)
		//Call Generate OTP
		otp, timeGenerate, errs := services.GenerateTOTP(dataBundleEncrypt, secretKey)
		if errs != nil {
			fmt.Println("Error generating OTP:", errs.Error())
			return
		}
		//otpCode, errs := util.GenerateOtpCustom(secretKey, time.Now())
		//if errs != nil {
		//	fmt.Printf("Error -> Cannot Generate OTP with error \n")
		//}
		timeOtp = timeGenerate
		//fmt.Println("Result Generate OTP : ", otp)
		fmt.Println(otp)
		//fmt.Println("Result Generated OTP : ", otpCode)

	//To Validate OTP call or run with (./generator validate <otp> <clientID>)
	case "validate":
		if len(os.Args) < 3 {
			fmt.Println("Usage: <dirPath> validate <otp> <clientID>")
			return
		}
		otp := os.Args[2]
		clientID = os.Args[3]
		//fmt.Println("action Command 0 : ", os.Args[0])
		//fmt.Println("action Command 1 : ", os.Args[1])
		//fmt.Println("action Command 2 : ", os.Args[2])
		//fmt.Println("action Command 3 : ", os.Args[3])
		// Get secret key for user1
		//secretKey, err := keyStore.GetSecretKey(clientID)
		//if err != nil {
		//	fmt.Println("Error retrieving secret key for user:", err.Error())
		//	return
		//}
		//fmt.Println("Secret key for user:", secretKey)

		valid, errs := services.Validat eTOTP(otp, timeOtp, secretKey)
		if errs != nil {
			fmt.Println("Error validating TOTP:", errs.Error())
			fmt.Println("Invalid OTP")
			return
		}
		//fmt.Println("Result Validate OTP : ", valid)
		fmt.Println(valid)

	//To Validate OTP call or run with (./generator createQr <dataBundle>)
	case "createQr":
		if len(os.Args) < 1 {
			fmt.Println("Usage: createQr <otp>")
			return
		}
		otp := os.Args[2]
		err := services.CreateQRCode(otp, "qrcode.png")
		if err != nil {
			fmt.Println("Error creating QR code:", err)
			return
		}
		fmt.Println("QR code created successfully.")

	default:
		fmt.Println("Invalid action:", action)
	}

	// Cleanup: Simulate expiration of OTP after a certain duration
	//time.Sleep(30 * time.Second)
	// Now, if you try to validate the same OTP again, it should not be valid.
}
