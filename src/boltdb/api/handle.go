package api

import (
	"boltdb/crud"
	"encoding/json"
	"net/http"

	"github.com/boltdb/bolt"
)

type Sayhi struct {
	id          uint64
	description string
}

func Add(db *bolt.DB) http.HandlerFunc {

	sayhi := crud.Insert_DB(db)

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		jsondata, _ := json.Marshal(sayhi)
		responseWriter.Write((jsondata))
	}
}

func Read(db *bolt.DB) http.HandlerFunc {

	Id, description := crud.Get_sayhi(db)
	jsonSayhi := Sayhi{
		id:          Id,
		description: description,
	}
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		jsondata, _ := json.Marshal(jsonSayhi)
		responseWriter.Write((jsondata))
	}
}
