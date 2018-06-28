package crud

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func ReadSayhi(db *leveldb.DB) string {

	defer db.Close()
	data, _ := db.Get([]byte("sayhi"), nil)
	fmt.Println(string(data))
	//	var sayhi model.Sayhi
	//json.Unmarshal([]byte(string(data)), &sayhi)
	//	fmt.Printf(" Description: %s\n", sayhi.Description)

	return string(data)

}
