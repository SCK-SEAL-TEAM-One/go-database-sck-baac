package api

import (
	"encoding/json"
	"net/http"
)

func GetReadAll(ReadAllKeyFunc func() ([]map[string]string, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		list, err := ReadAllKeyFunc()
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		jsonResponse, _ := json.Marshal(list)
		responseWriter.Write(jsonResponse)

	}
}

func GetReadById(ReadByIdFunc func(string) (map[string]string, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		queryString := request.URL.Query()
		id := queryString.Get("id")
		list, err := ReadByIdFunc(id)
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		jsonResponse, _ := json.Marshal(list)
		responseWriter.Write(jsonResponse)

	}
}
