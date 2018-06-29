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
		record, err := ReadByIdFunc(id)
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return

		}
		jsonResponse, _ := json.Marshal(record)
		responseWriter.Write(jsonResponse)

	}
}
func GetCreate(AddFunc func(string) (map[string]string, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}

		record, err := AddFunc(postData["description"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}
		jsonResponse, _ := json.Marshal(record)
		responseWriter.Write(jsonResponse)
	}
}

func GetEdit(EditFunc func(string, string) (map[string]string, error)) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		queryString := request.URL.Query()
		id := queryString.Get("id")
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}

		record, err := EditFunc(id, postData["description"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)
			return
		}
		jsonResponse, _ := json.Marshal(record)
		responseWriter.Write(jsonResponse)
	}
}
