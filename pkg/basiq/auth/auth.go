package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/koustubh25/babybucks/pkg/basiq/util"
)

// this function expects the API key to be passed using an environment variable `BASIQ_APIKEY`
func Get_client_access_token(userId string) util.Client_token_response {
	url := util.BASE_ENDPOINT + RELATIVE_TOKEN_PATH
	payload := strings.NewReader("scope=CLIENT_ACCESS&userId=" + userId)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic "+os.Getenv("BASIQ_APIKEY"))
	req.Header.Add("basiq-version", "3.0")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var client_token_response util.Client_token_response
	if json.Unmarshal([]byte(string(body)), &client_token_response) != nil {
		panic(err)
	}
	return client_token_response
}
