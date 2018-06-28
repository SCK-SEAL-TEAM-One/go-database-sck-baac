package main

import (
	"boltdb/crud"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}

	// Open the database.
	db, err := bolt.Open("src/boltdb/test.db", 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/setperson", func(w http.ResponseWriter, r http.Request){
		u := r.URL.Query()
		crud.Insert_DB(db)
	}
	crud.View_DB(db)

	// Close database to release the file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
	
	http.ListenAndServe("localhost:8000",nil)
}
