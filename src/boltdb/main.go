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

	http.HandleFunc("/read", api.Add(db))
	http.ListenAndServe(":3000", nil)
}
