package crud

import (
	"couchdb/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ReadCouchDB(id string, couchdbURL string) model.Docs {
	url := couchdbURL + "/sckseal/_find"
	couchDBResponse := model.Docs{}
	payload := strings.NewReader("{\"selector\": {\"id\": " + id + "},\"fields\":[\"id\",\"sayhi\"]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))

	err := json.Unmarshal(body, &couchDBResponse)
	if err != nil {
		fmt.Println("error : ", err)
	}
	return couchDBResponse
}
