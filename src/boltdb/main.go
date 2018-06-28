package main

import (
	"boltdb/crud"
)

func main() {

	crud.Connect()
	crud.Insert_DB()
	// crud.View_DB()
}
