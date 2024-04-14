package services

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

// CreateQRCode creates a QR code for the given data and saves it to a file
func CreateQRCode(data string, filename string) error {
	err := qrcode.WriteFile(data, qrcode.Medium, 256, filename)
	if err != nil {
		return err
	}
	fmt.Println("QR code has been saved to", filename)

	return nil
}

//qrData, err := json.Marshal(g)
//if err != nil {
//return
//}

//--- Prepare Data
//qrDataValue := flag.String("data", string(qrData), "Data For QR Code")
//qrOutputValue := flag.String("output", "qrcode.png", "File QR Code")

//flag.Parse()

//--- Create QR code
//err = qrcode.WriteFile(*qrDataValue, qrcode.Medium, 256, *qrOutputValue)
//if err != nil {
//fmt.Printf("Error -> Cannot Generate QR Code")
//return
//}

//func produceQR(grochatBundle DtoModel.RequestGroChatBundle) (err error) {
//	var (
//		gro        = grochatBundle.DataDetail
//		secretTemp = gro.ClientID + gro.HwID + gro.TimestampStr + gro.Type
//		key        *otp.Key
//	)
//
//	//Account Name
//	//data yg di enkripsi bundle secretTemp
//	gro.Timestamp, err = util.TimeParse(constanta.TimeLayoutDefault, gro.TimestampStr)
//	if err != nil {
//		return
//	}
//
//	key, err = GenerateTOTP(secretTemp)
//	if err != nil {
//		return
//	}
//
//	fmt.Println(key)
//	return
//}
