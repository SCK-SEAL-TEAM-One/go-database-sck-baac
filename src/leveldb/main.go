package main

import (
	"leveldb/api"
	"log"
	"net/http"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {

	db, err := leveldb.OpenFile("/Users/captainamerica/Documents/ldb/sckseal", nil)
	if err != nil {
		log.Fatal("connect error!")
	}
	http.HandleFunc("/helloworld/readAll", api.ReadAll(db))
	http.HandleFunc("/helloworld/read", api.Read(db))
	http.ListenAndServe(":8080", nil)
}
