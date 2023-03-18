package main

import (
	influx "github.com/koustubh25/babybucks/internal/influxdb"
)

func main() {

	// client_access_token := auth.Get_client_access_token(os.Getenv("BASIC_USERID")).Access_token
	// a := accounts.Get_accounts(os.Getenv("BASIC_USERID"), client_access_token)[0]

	influx.GenerateMockData()

}
