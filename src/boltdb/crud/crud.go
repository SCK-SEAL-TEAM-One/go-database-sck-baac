package crud

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func Insert_DB(db *bolt.DB) {
	// Execute several commands within a read-write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("widgets"))
		b.Put([]byte("Name"), []byte("Lek"))
		b.Put([]byte("Age"), []byte("22"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}

func View_DB(db *bolt.DB) {
	// Read the value back from a separate read-only transaction.
	if err := db.View(func(tx *bolt.Tx) error {
		value := tx.Bucket([]byte("widgets")).Get([]byte("foo"))
		fmt.Printf("The value of 'foo' is: %s\n", value)
		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
