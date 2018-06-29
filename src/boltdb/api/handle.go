package api

import (
	"boltdb/crud"
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
)

func Add(db *bolt.DB) http.HandlerFunc {

	sayhi := crud.Insert_DB(db)

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		jsondata, _ := json.Marshal(sayhi)
		responseWriter.Write((jsondata))
	}
}
