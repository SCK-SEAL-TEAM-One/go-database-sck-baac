package crud

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func ReadCouchDB(id int, couchdbIP string) {
	url := couchdbIP + "/sckseal/_find"

	payload := strings.NewReader("{\"selector\": {\"id\": " + strconv.Itoa(id) + "}}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	//return nil
}
