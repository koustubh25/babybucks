package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://au-api.basiq.io/users/userId/accounts"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
