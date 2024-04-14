package dtoModel

type ResponseGenerateOTP struct {
	OTP string `json:"otp"`
}

type ResponseValidateOTP struct {
	ValidStatus bool `json:"valid_status"`
}

type ResponseQrCode struct {
	QrCode string `json:"qr_code"`
}
