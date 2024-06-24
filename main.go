package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"totp-learn/services"
	"totp-learn/util"
)

func main() {
	var (
		err error
		//dataBundle dtoModel.RequestDataBundleOtp
		secretKey string
		timeOtp   time.Time
		//clientID   string
	)

	// Command to execute in Linux environment using WSL
	//cmd := "wsl"
	//args := []string{"/home/martinus/generator_time", "generate", "b934d7e43a3b0f36ab442544e4055de96d57e77bfccafc1be912"}

	// Create command
	//command := exec.Command(cmd, args...)

	// Run command and capture output
	//output, errs := command.CombinedOutput()
	//output, errs := command.Output()
	//if errs != nil {
	//	fmt.Println("Error executing command:", errs.Error())
	//	return
	//}

	// Print output
	//fmt.Println("Output from Linux environment:", string(output))

	//fmt.Println("Encode Data Bundle : ", util.EncodeSHA256("47e8d0327c1a43cfa1f83195aa60db66_B-/1HRD8T2/CNCMC0092J00A7/ - Audentis Fortuna iuvat_2024-04-09 15:57:50.903_ND6-sysadminLv1"))
	//fmt.Println("Encode Data Bundle : ", util.EncodeSHA256("47e8d0327c1a43cfa1f83195aa60db66_B-/1HRD8T2/CNCMC0092J00A7/ - Audentis Fortuna iuvat_2024-04-16 15:30:40.702_ND6-sysadminLv1"))
	//fmt.Println("Encode Data Bundle : ", util.CheckSumWithSha256([]byte("47e8d0327c1a43cfa1f83195aa60db66_B-/1HRD8T2/CNCMC0092J00A7/ - Audentis Fortuna iuvat_2024-04-09 15:57:50.903_ND6-sysadminLv1")))

	// Initialize SecretKeyStore
	//keyStore := services.NewSecretKeyStore()
	// The key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	secretKey = "0123456789abcdef0123456789abcdef"

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

	if len(os.Args) < 2 {
		log.Fatal("Arguments Empty")
	}
	action := os.Args[1]

	switch action {
	//To Generate OTP call or run with (./generator generate <dataBundle>)
	case "generate":
		if len(os.Args) < 2 || os.Args[1] != "generate" {
			fmt.Println("Usage: <dirPath> generate <dataBundle>")
			return
		}
		if len(os.Args) < 3 {
			fmt.Println("Data that want to Generate empty!")
			return
		}
		dataBundleEncrypt := os.Args[2]

		//Decode secretKey
		//dataChipper, err := util.DecodeSHA256(dataBundleEncrypt)
		//if err != nil {
		//	fmt.Println("Error Decode Data Bundle :", err.Error())
		//	return
		//}

		//util.CheckSumWithSha256([]byte(dataBundleEncrypt))

		//Generate secretKey
		secretKey, err = util.GenerateOtpSecretKey([]byte(dataBundleEncrypt))

		//Decode Data
		//decodedData, err := base64.StdEncoding.DecodeString(dataBundleEncrypt)
		//if err != nil {
		//	fmt.Println("Decoding error : ", err)
		//	return
		//}
		//secretKeyByte := []byte(secretKey)
		//result, errs := util.DecryptData(decodedData, secretKeyByte)
		//if errs != nil {
		//	fmt.Println("Error Decrypt data : ", err)
		//	return
		//}
		//resString := util.ByteArrayToString(result)
		//dataSplit := strings.Split(resString, "_")

		//Store data bundle in struct
		//dataCheckSum := util.CheckSumWithSha256(dataChipper)
		//dataBundleString := util.ByteArrayToString(dataChipper)
		//fmt.Println("Data bundle string = ", dataBundleEncrypt)
		//dataSplit := strings.Split(dataBundleEncrypt, "_")
		//dataRequest := dataBundle.DataDetail
		//if len(dataSplit) < 2 {
		//	fmt.Println("Error Data Bundle Invalid!")
		//	return
		//}
		//if len(dataSplit) > 3 {
		//	dataRequest.ClientID = dataSplit[0]
		//	dataRequest.HwID = dataSplit[1]
		//	dataRequest.TimestampStr = dataSplit[2]
		//	dataRequest.Type = dataSplit[3]
		//}
		//Store clientId into global var
		//clientID = dataRequest.ClientID

		//Store secretKey
		//keyStore.SetSecretKey(clientID, secretKey)
		//secretKey, errs := keyStore.GetSecretKey(clientID)
		//if errs != nil {
		//	fmt.Println("Error retrieving secret key for user : ", errs.Error())
		//	return
		//}
		//fmt.Println("Secret key for user:", secretKey)
		//Call Generate OTP
		otp, timeGenerate, errs := services.GenerateTOTP(dataBundleEncrypt, secretKey)
		if errs != nil {
			fmt.Println("Error generating OTP : ", errs.Error())
			return
		}
		timeOtp = timeGenerate
		fmt.Println(otp)

	//To Validate OTP call or run with (./generator validate <otp> <dataBundle>)
	case "validate":
		if len(os.Args) < 3 || os.Args[1] != "validate" {
			fmt.Println("Usage: <dirPath> validate <otp> <dataBundle>")
			return
		}
		otp := os.Args[2]
		if otp == "" {
			fmt.Println("OTP empty!")
			return
		}
		if len(os.Args) < 4 {
			fmt.Println("Data Validate empty!")
			return
		}
		dataBundleEncrypt := os.Args[3]

		// Get secret key for user1
		//secretKey, err = keyStore.GetSecretKey(clientID)
		//if err != nil {
		//	fmt.Println("Error retrieving secret key for user : ", err.Error())
		//	return
		//}
		secretKey, err = util.GenerateOtpSecretKey([]byte(dataBundleEncrypt))
		//fmt.Println("Secret key for user:", secretKey)

		valid, errs := services.ValidateTOTP(otp, timeOtp, secretKey)
		if errs != nil {
			fmt.Println("Invalid OTP :", errs.Error())
			//fmt.Println("Invalid OTP")
			return
		}
		fmt.Println(valid)

	//To Validate OTP call or run with (./generator createQr <dataBundle>)
	case "createQr":
		if len(os.Args) < 1 {
			fmt.Println("Usage: <dirPath> createQr <dataBundle>")
			return
		}
		otp := os.Args[2]
		err := services.CreateQRCode(otp, "qrcode.png")
		if err != nil {
			fmt.Println("Error creating QR code : ", err)
			return
		}
		fmt.Println("QR code created successfully.")

	//case "encodeData":
	//	if len(os.Args) < 2 {
	//		fmt.Println("Usage: <dirPath> encodeData <dataBundle>")
	//		return
	//	}
	//	dataBundlePlain := os.Args[2]
	//	dataBundlePlainByte := []byte(dataBundlePlain)
	//	secretKeyByte := []byte(secretKey)
	//	result, errs := util.EncryptData(dataBundlePlainByte, secretKeyByte)
	//	if errs != nil {
	//		fmt.Println("Error encode data : ", err)
	//		return
	//	}
	//	fmt.Println(base64.StdEncoding.EncodeToString(result))

	//case "decodeData":
	//	if len(os.Args) < 2 {
	//		fmt.Println("Usage: <dirPath> decodeData <dataBundle>")
	//		return
	//	}
	//	dataBundleEncrypt := os.Args[2]
	//	// Decode base64-encoded string
	//	decodedData, err := base64.StdEncoding.DecodeString(dataBundleEncrypt)
	//	if err != nil {
	//		fmt.Println("Decoding error:", err)
	//		return
	//	}
	//	secretKeyByte := []byte(secretKey)
	//	result, errs := util.DecryptData(decodedData, secretKeyByte)
	//	if errs != nil {
	//		fmt.Println("Error decode data : ", err)
	//		return
	//	}
	//	resString := util.ByteArrayToString(result)
	//	fmt.Println(resString)

	default:
		fmt.Println("Invalid action:", action)
	}

	// Cleanup: Simulate expiration of OTP after a certain duration
	//time.Sleep(30 * time.Second)
	// Now, if you try to validate the same OTP again, it should not be valid.
}
