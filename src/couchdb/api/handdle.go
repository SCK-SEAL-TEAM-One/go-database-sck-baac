package api

import (
	"couchdb/crud"
	"encoding/json"
	"net/http"
)

const (
	DBUSERNAME = "sckseal"
	DBPASSWORD = "123456"
	DBIP       = "127.0.0.1:5984"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	description := r.URL.Query().Get("description")
	couchDBResponse := crud.AddCouchDB(DBUSERNAME, DBPASSWORD, DBIP, id, description)
	jsonCouchDBResponse := []byte(couchDBResponse)
	w.Write(jsonCouchDBResponse)
}

func ReadHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	couchDBResponse := crud.ReadCouchDB(DBIP, id)
	jsonCouchDBResponse, _ := json.Marshal(couchDBResponse)
	w.Write(jsonCouchDBResponse)
}

func EditHandeler(w http.ResponseWriter, r *http.Request) {
	couchDbId := r.URL.Query().Get("couchdbid")
	id := r.URL.Query().Get("id")
	description := r.URL.Query().Get("description")
	revision := r.URL.Query().Get("revision")
	couchDBResponse := crud.EditCouchDB(DBUSERNAME, DBPASSWORD, DBIP, couchDbId, id, description, revision)
	jsonCouchDBResponse, _ := json.Marshal(couchDBResponse)
	w.Write(jsonCouchDBResponse)
}

func DeleteHandeler(w http.ResponseWriter, r *http.Request) {
	couchDbId := r.URL.Query().Get("couchdbid")
	revision := r.URL.Query().Get("revision")
	couchDBResponse := crud.DeleteCouchDB(DBIP, couchDbId, revision)
	jsonCouchDBResponse, _ := json.Marshal(couchDBResponse)
	w.Write(jsonCouchDBResponse)
}
