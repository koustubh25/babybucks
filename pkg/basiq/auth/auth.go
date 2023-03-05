package auth

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// this function expects the API key to be passed using an environment variable `API_KEY`
func Get_client_access_token() {
	url := BASE

	req, _ := http.NewRequest("POST", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
