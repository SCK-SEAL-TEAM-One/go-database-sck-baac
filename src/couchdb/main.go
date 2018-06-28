package main

import (
	"couchdb/api"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/read", api.ReadHandler)
	fmt.Println("Listening on port 9092")
	http.ListenAndServe(":9092", nil)
}
