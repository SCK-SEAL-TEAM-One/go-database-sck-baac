package crud

import (
	"encoding/binary"
	"log"

	"github.com/boltdb/bolt"
)

type Sayhi struct {
	id          uint64
	description string
}

func Insert_DB(db *bolt.DB) string {
	// Execute several commands within a read-write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("sayhi"))
		b.Put([]byte("id"), []byte("1"))
		b.Put([]byte("description"), []byte("Hello World"))
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return "True"
}

func Get_sayhi(db *bolt.DB) (uint64, string) {
	// jsonSayhi := Sayhi{}
	var description string
	var idToInt uint64
	if err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sayhi"))
		id := b.Get([]byte("id"))
		description = string(b.Get([]byte("description")))
		idToInt = binary.BigEndian.Uint64(id)

		// jsonSayhi = Sayhi{
		// 	id:          idToInt,
		// 	description: string(description),
		// }

		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return idToInt, description
}
