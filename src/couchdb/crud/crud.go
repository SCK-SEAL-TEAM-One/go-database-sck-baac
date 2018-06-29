package crud

import (
	"couchdb/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func AddCouchDB(dbUsername, dbPassword, dbIP, id, description string) string {
	couchDbId := generateCouchDbUuids(dbIP)
	url := "http://" + dbUsername + ":" + dbPassword + "@" + dbIP + "/sckseal/" + couchDbId + "/"
	postRequestData := `{"Id": ` + id + `,"Description": "` + description + `"}`
	payload := strings.NewReader(postRequestData)
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func generateCouchDbUuids(dbIP string) string {
	url := "http://" + dbIP + "/_uuids"
	couchDbId := model.CouchDbId{}
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &couchDbId)
	if err != nil {
		fmt.Println("error : ", err)
	}
	return couchDbId.Uuids[0]
}

func ReadCouchDB(dbIP, id string) model.Docs {
	url := "http://" + dbIP + "/sckseal/_find"
	couchDbResponse := model.Docs{}
	postRequestData := `{"selector": {"Id": ` + id + `}}`
	payload := strings.NewReader(postRequestData)
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal(body, &couchDbResponse)
	if err != nil {
		fmt.Println("error : ", err)
	}
	return couchDbResponse
}

func EditCouchDB(dbUsername, dbPassword, dbIP, couchDbId, id, description, revision string) string {
	url := "http://" + dbUsername + ":" + dbPassword + "@" + dbIP + "/sckseal/" + couchDbId + "/"
	postRequestData := `{"Id":` + id + `,"Description" : "` + description + `" , "_rev" : "` + revision + `"}`
	payload := strings.NewReader(postRequestData)
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func DeleteCouchDB(dbIP, couchDbId, revision string) string {
	url := "http://" + dbIP + "/sckseal/" + couchDbId + "?rev=" + revision
	req, _ := http.NewRequest("DELETE", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
