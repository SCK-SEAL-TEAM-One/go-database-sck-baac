package crud

import (
	"encoding/json"
	"fmt"
	"leveldb/model"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

func ConnectLevelDB() *leveldb.DB {
	db, err := leveldb.OpenFile("/Users/captainamerica/Documents/ldb/sckseal", nil)
	if err != nil {
		log.Fatal("connect error!")
	}

	return db
}

func ReadSayhi() model.Sayhi {

	db := ConnectLevelDB()
	defer db.Close()
	data, _ := db.Get([]byte("sayhi"), nil)

	var sayhi model.Sayhi
	json.Unmarshal([]byte(string(data)), &sayhi)
	fmt.Printf(" Description: %s\n", sayhi.Description)

	return sayhi

}
