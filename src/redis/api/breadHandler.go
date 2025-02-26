package api

import (
	"encoding/json"
	"net/http"
)

func GetRead(ReadKeyFunc func(string) (string, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		queryString := request.URL.Query()
		key := queryString.Get("key")
		val, err := ReadKeyFunc(key)
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		responseWriter.Write([]byte(val))

	}
}

func GetBreadCreate(WriteKeyFunc func(string, interface{}) error) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}
		err = WriteKeyFunc(postData["key"], postData["value"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}
		responseWriter.Write([]byte("Create Successful"))
	}
}

func GetDelete(DeleteKeyFunc func(string) (int64, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		_, err = DeleteKeyFunc(postData["key"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		responseWriter.Write([]byte("Delete Successful"))
	}
}
