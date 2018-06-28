package main

import (
	"api"
	"net/http"
)

func main() {
	//crud.ConnectLevelDB("/Users/captainamerica/Documents/ldb/sckseal")
	http.HandleFunc("/read", api.Read)
	http.ListenAndServe(":8080", nil)
}
