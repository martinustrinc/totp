package services

import (
	"errors"
	"fmt"
	"time"
	"totp-learn/model"
	"totp-learn/util"
)

// ValidateTOTP checks if the input meets certain criteria (dummy validation for demonstration)
func ValidateTOTP(inputOtp string, timeOtp time.Time, secretKey string) (result string, err error) {
	var dataBundle dtoModel.RequestDataBundleOtp
	//fmt.Println("Length OTP : ", len(inputOtp))
	if len(inputOtp) < 6 {
		fmt.Println("OTP must be 6 characters")
		return result, errors.New("input must be 6 characters")
	}

	//format data bundle (dataRequest.ClientID + "_" + dataRequest.HwID + "_" + dataRequest.TimestampStr + "_" + dataRequest.Type)
	dataRequest := dataBundle.DataDetail
	//dataRequest.Timestamp, err = util.TimeParse(time.RFC3339, dataRequest.TimestampStr)
	dataRequest.Timestamp = timeOtp
	if err != nil {
		return result, err
	}
	//dataValue := dataRequest.ClientID + "_" + dataRequest.HwID + "_" + dataRequest.TimestampStr + "_" + dataRequest.Type

	// Generate a new TOTP key
	// key, err := totp.Generate(totp.GenerateOpts{
	//	Issuer:      "NexSOFT-TOTP",
	//	AccountName: "nextrac@nexcloud.co.id",
	//	Algorithm:   otp.AlgorithmSHA256,
	//	Digits:      otp.DigitsSix,
	//	Secret:      []byte(util.CheckSumWithSha256([]byte(dataValue))),
	//	SecretSize:  20,
	//})
	//if err != nil {
	//	fmt.Printf("Error -> Cannot Generate Key with error \n")
	//	return result, err
	//}

	//secretKey := key.Secret()
	key := util.StringToBase32String(secretKey)
	timeOtp = time.Now()
	valid, errs := util.ValidateOtpCustom(key, inputOtp, timeOtp)
	//valid := util.ValidateOtpBasic(key, inputOtp)
	if errs != nil {
		//fmt.Println("Error Validate OTP : ", errs.Error())
	}
	if valid {
		result = "Success Validate OTP"
	} else {
		result = "Failed Validate OTP"
	}

	//fmt.Println("Time Generate : ", dataRequest.Timestamp)
	//fmt.Println("Time Generate Unix Milli : ", dataRequest.Timestamp.UnixMilli())
	//fmt.Println("Time Validate : ", timeOtp)
	//fmt.Println("Time Validate Unix Milli : ", timeOtp.UnixMilli())
	if timeOtp.UnixMilli()-dataRequest.Timestamp.UnixMilli() <= 0 {
		result = "Failed Validate OTP, OTP Expired"
	}

	//if needed to response JSON
	//resp, err := json.Marshal(result)
	//if err != nil {
	//	return result, err
	//}
	//fmt.Println("Response Validate OTP Json -> ", resp)

	return result, err
}
