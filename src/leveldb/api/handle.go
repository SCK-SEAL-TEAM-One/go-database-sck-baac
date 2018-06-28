package api

import (
	"encoding/json"
	"leveldb/crud"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

// type Connect struct {
// 	DB *leveldb.DB
// }

func Read(db *leveldb.DB) http.HandlerFunc {

	sayhi := crud.ReadSayhi(db)

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}
