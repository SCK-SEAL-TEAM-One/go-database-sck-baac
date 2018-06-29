package crud

import (
	"log"

	"github.com/boltdb/bolt"
)

func Insert_DB(db *bolt.DB) string {
	// Execute several commands within a read-write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("widgets"))
		b.Put([]byte("Name"), []byte("Lek"))
		b.Put([]byte("Age"), []byte("22"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return "True"
}
