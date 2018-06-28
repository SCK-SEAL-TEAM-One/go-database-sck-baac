package api

import (
	"couchdb/crud"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	URL = "http://127.0.0.1:5984"
)

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	querySayhiByID := r.URL.Query().Get("id")
	fmt.Println("id : " + querySayhiByID)
	couchDBResponse := crud.ReadCouchDB(querySayhiByID, URL)
	jsonCouchDBResponse, _ := json.Marshal(couchDBResponse)
	w.Write(jsonCouchDBResponse)
}
