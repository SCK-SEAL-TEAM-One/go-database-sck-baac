package main

import (
	"couchdb/api"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/add", api.AddHandler)
	http.HandleFunc("/read", api.ReadHandler)
	http.HandleFunc("/edit", api.EditHandeler)
	http.HandleFunc("/delete", api.DeleteHandeler)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
