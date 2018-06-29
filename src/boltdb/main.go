package main

import (
	"boltdb/api"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("src/boltdb/sckseal.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/add", api.Add(db))
	http.HandleFunc("/read", api.Read(db))
	http.ListenAndServe(":3000", nil)
}
