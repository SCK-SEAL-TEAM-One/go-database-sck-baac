package main

import (
	"encoding/json"
	"net/http"
	"redis/crud"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	rc := crud.RedisClient{Client: client}

	http.HandleFunc("/redis/read", func(responseWriter http.ResponseWriter, request *http.Request) {
		queryString := request.URL.Query()
		key := queryString.Get("key")
		val, err := rc.ReadKey(key)
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)

		}
		responseWriter.Write([]byte(val))

	})

	http.HandleFunc("/redis/create", func(responseWriter http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)

		}
		err = rc.WriteKey(postData["key"], postData["value"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)

		}
		responseWriter.Write([]byte("Create Successful"))
	})

	http.HandleFunc("/redis/delete", func(responseWriter http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var postData map[string]string
		err := decoder.Decode(&postData)

		if err != nil {
			http.Error(responseWriter, err.Error(), 500)

		}
		val, err := rc.DeleteKey(postData["key"])
		if err != nil {
			http.Error(responseWriter, err.Error(), 500)

		}
		responseWriter.Write([]byte("Delete Successful" + string(val)))
	})

	http.ListenAndServe(":3000", nil)

}
