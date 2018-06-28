package main

import (
	"couchdb/crud"
)

const (
	url = "http://127.0.0.1:5984"
)

func main() {
	crud.ReadCouchDB(2, url)
}
