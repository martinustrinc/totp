package domain

import "time"

type RequestGrochatBundle struct {
	GrochatAccount string                  `json:"grochat_account"`
	DataDetail     RequestDetailDataBundle `json:"data"`
}

type RequestDetailDataBundle struct {
	ClientID     string `json:"client_id"`
	HWID         string `json:"hwid"`
	Type         string `json:"type"`
	TimestampStr string `json:"timestamp"`
	Timestamp    time.Time
}
