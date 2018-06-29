package api

import (
	"encoding/json"
	"net/http"
	"postgresql/crud"
)

func Read() http.HandlerFunc {

	sayhi := crud.ReadSayhi()

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}

func ReadById() http.HandlerFunc {

	sayhi := crud.ReadSayhi()

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}

func Add() http.HandlerFunc {

	sayhi := crud.ReadSayhi()

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}

func Edit() http.HandlerFunc {

	sayhi := crud.ReadSayhi()

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}

func Delete() http.HandlerFunc {

	sayhi := crud.ReadSayhi()

	return func(responseWriter http.ResponseWriter, request *http.Request) {
		JSONConnectData, _ := json.Marshal(sayhi)
		responseWriter.Write(JSONConnectData)
	}
}
