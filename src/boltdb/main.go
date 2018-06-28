package main

import (
	"boltdb/crud"
	"encoding/json"
	"log"
	"net/http"

	"github.com/boltdb/bolt"
)

func main() {
	type Person struct {
		Name string
		Age  string
	}

	// Open the database.
	db, _ := bolt.Open("src/boltdb/test.db", 0666, nil)

	http.HandleFunc("/setperson", func(w http.ResponseWriter, r *http.Request) {
		crud.Insert_DB(db)
	})

	http.HandleFunc("/getperson", func(w http.ResponseWriter, r *http.Request) {

		// Read the value back from a separate read-only transaction.
		db.View(func(tx *bolt.Tx) error {
			widgets := tx.Bucket([]byte("widgets"))
			person := Person{
				Name: string(widgets.Get([]byte("Name"))),
				Age:  string(widgets.Get([]byte("Age"))),
			}
			json.NewEncoder(w).Encode(&person)
			return nil
		})
	})
	// Close database to release the file lock.
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe("localhost:8000", nil)
}

//sfunc Read(db *bolt.DB) ht
