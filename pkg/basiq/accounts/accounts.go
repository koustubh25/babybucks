package accounts

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/koustubh25/babybucks/pkg/basiq/users"
	"github.com/koustubh25/babybucks/pkg/basiq/util"
	"github.com/tidwall/gjson"
)

func Get_accounts(userId string, bearerToken string) util.Accounts {

	url := util.BASE_ENDPOINT + users.RELATIVE_USERS_PATH + "/" + userId + RELATIVE_ACCOUNTS_PATH
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+bearerToken)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var accounts util.Accounts
	a := gjson.Get(string(body), "data")
	if json.Unmarshal([]byte(a.String()), &accounts) != nil {
		panic(err)
	}
	return accounts
}
