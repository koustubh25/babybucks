package util

const BASE_ENDPOINT = "https://au-api.basiq.io"

type Client_token_response struct {
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	Expires_in   int32  `json:"expires_in"`
}

type Accounts []Account

type Account struct {
	Account_holder_name string  `json:"accountHolder"`
	Balance             float32 `json:"balance,string"`
	Available_funds     float32 `json:"availableFunds,string"`
	Id                  string  `json:"id"`
}
