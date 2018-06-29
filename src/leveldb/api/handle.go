package api

import (
	"encoding/json"
	"leveldb/crud"
	"leveldb/model"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

// type Connect struct {
// 	DB *leveldb.DB
// }

func ReadAll(db *leveldb.DB) http.HandlerFunc {

	sayhi := crud.ReadSayhi(db)

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}

func Read(db *leveldb.DB) http.HandlerFunc {
	var sayhi model.Sayhi = crud.ReadSayhi(db)

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}
