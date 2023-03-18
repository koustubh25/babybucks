package consent

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/koustubh25/babybucks/pkg/basiq/users"
	"github.com/koustubh25/babybucks/pkg/basiq/util"
)

func Get_consents(userId string, bearerToken string) {

	url := util.BASE_ENDPOINT + users.RELATIVE_USERS_PATH + "/" + userId + RELATIVE_CONSENTS_PATH

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer "+bearerToken)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
