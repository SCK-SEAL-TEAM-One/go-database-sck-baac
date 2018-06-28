package api

import (
	"encoding/json"
	"leveldb/crud"
	"leveldb/model"
	"net/http"
)

type Connect struct {
	DB string
	Id string
}

func Read(responseWriter http.ResponseWriter, request *http.Request) {

	var sayhi model.Sayhi = crud.ReadSayhi()

	JSONConnectData, _ := json.Marshal(sayhi)

	responseWriter.Write(JSONConnectData)
}
