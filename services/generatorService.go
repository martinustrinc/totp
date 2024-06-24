package services

import (
	"time"
	"totp-learn/model"
	"totp-learn/util"
)

var (
//timeLayout = constanta.TimeLayoutDefault
)

// GenerateTOTP generates a Time-Based One-Time Password
func GenerateTOTP(dataEncryption string, secretKey string) (result string, timeGenerate time.Time, err error) {
	var (
		otpCode string
		//key        *otp.Key
		dataBundle dtoModel.RequestDataBundleOtp
	)

	//format data bundle (dataRequest.ClientID + "_" + dataRequest.HwID + "_" + dataRequest.TimestampStr + "_" + dataRequest.Type)
	//encode data bundle into string
	//dataInput, err := util.DecodeSHA256(dataEncryption)
	//dataString := util.ByteArrayToString(dataInput)
	//dataString := dataEncryption
	//dataSplit := strings.Split(dataString, "_")
	dataRequest := dataBundle.DataDetail
	//if len(dataSplit) < 2 {
	//	fmt.Println("Error length")
	//	return
	//}
	//if len(dataSplit) > 3 {
	//	dataRequest.ClientID = dataSplit[0]
	//	dataRequest.HwID = dataSplit[1]
	//dataRequest.TimestampStr = dataSplit[2]
	//dataRequest.TimestampStr = time.Now().String()
	dataRequest.Timestamp = time.Now()
	//dataRequest.Type = dataSplit[3]
	//}
	//dataRequest.Timestamp, err = util.TimeParse(timeLayout, dataRequest.TimestampStr)
	//if err != nil {
	//	fmt.Println("Error Parsing")
	//	return
	//}

	// Generate a new TOTP key
	//key, err = totp.Generate(totp.GenerateOpts{
	//	Issuer:      "NexSOFT-TOTP",
	//	AccountName: "nextrac@nexcloud.co.id",
	//	Algorithm:   otp.AlgorithmSHA256,
	//	Digits:      otp.DigitsSix,
	//	Secret:      []byte(util.CheckSumWithSha256(dataInput)),
	//	SecretSize:  20,
	//})
	//if err != nil {
	//	fmt.Printf("Error -> Cannot Generate Key with error \n")
	//	return result, err
	//}

	//fmt.Println("Masuk Generator Service")
	//secretKey := key.Secret()

	key := util.StringToBase32String(secretKey)
	// Generate TOTP
	otpCode, err = util.GenerateOtpCustom(key, dataRequest.Timestamp)
	if err != nil {
		//fmt.Printf("Error -> Cannot Generate OTP with error \n")
		return
	}
	//fmt.Println("OTP Code -> ", otpCode)
	//responseGenerateOtp := dtoModel.ResponseGenerateOTP{
	//	OTP: otpCode,
	//}

	//if needed to response JSON
	//resp, err := json.Marshal(responseGenerateOtp)
	//if err != nil {
	//	return
	//}
	//fmt.Println("Response Generate OTP Json -> ", resp)

	result = otpCode
	timeGenerate = dataRequest.Timestamp
	return
}
