package dtoModel

import "time"

type RequestDataBundleOtp struct {
	UserAccount string                  `json:"user_account"`
	DataDetail  RequestDetailDataBundle `json:"data_encrypt"`
}

type RequestDetailDataBundle struct {
	UserAccount  string `json:"user_account"`
	ClientID     string `json:"client_id"`
	HwID         string `json:"hw_id"`
	Type         string `json:"type"`
	TimestampStr string `json:"timestamp"`
	Timestamp    time.Time
}
