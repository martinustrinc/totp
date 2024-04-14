package validation

import (
	"strings"
	"time"
	dtoModel "totp-learn/model"
	"totp-learn/util"
)

func ValidateDataBundle(dataValue []byte) (result string, err error) {
	var (
		dataBundle dtoModel.RequestDataBundleOtp
		//timeLayout = constanta.TimeLayoutDefault
	)

	dataString := util.ByteArrayToString(dataValue)
	dataSplit := strings.Split("_", dataString)
	dataRequest := dataBundle.DataDetail
	if len(dataSplit) < 3 {
		return
	}
	if len(dataSplit) > 3 {
		dataRequest.ClientID = dataSplit[0]
		dataRequest.HwID = dataSplit[1]
		dataRequest.TimestampStr = dataSplit[2]
		dataRequest.Type = dataSplit[3]
	}
	dataRequest.Timestamp, err = util.TimeParse(time.RFC3339, dataRequest.TimestampStr)
	if err != nil {
		return
	}

	result = dataRequest.ClientID + "_" + dataRequest.HwID + "_" + dataRequest.TimestampStr + "_" + dataRequest.Type
	return
}
