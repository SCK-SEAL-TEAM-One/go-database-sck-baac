package crud

import (
	"encoding/json"
	"fmt"
	"leveldb/model"

	"github.com/syndtr/goleveldb/leveldb"
)

func ReadSayhiAll(db *leveldb.DB) []model.Sayhi {

	defer db.Close()
	// data, _ := db.Get([]byte("sayhi"), nil)
	// fmt.Println(string(data))

	// return string(data)

	//listSayhi := make([]model.Sayhi, 0)
	data, _ := db.Get([]byte("sayhi"), nil)

	//aa := string(data)
	var sayhi []model.Sayhi
	json.Unmarshal([]byte(string(data)), &sayhi)

	//json.Unmarshal(data, &listSayhi)

	// iter := db.NewIterator(nil, nil)
	// for iter.Next() {
	// 	key := iter.Key()
	// 	value := iter.Value()
	// 	fmt.Printf("key: %s | value: %s\n", key, value)
	// 	sayhi := model.Sayhi{Id: string(key), Description: string(value)}
	// 	fmt.Println(sayhi)
	// 	listSayhi.List = append(listSayhi.List, sayhi)

	// }
	//sayhiData := model.ListSayhi{List: listSayhi}
	fmt.Println(sayhi)
	return sayhi
}

func ReadSayhi(db *leveldb.DB) model.Sayhi {

	defer db.Close()
	data, _ := db.Get([]byte("sayhi2"), nil)

	var sayhi model.Sayhi
	json.Unmarshal([]byte(string(data)), &sayhi)
	fmt.Printf(" Description: %s\n", sayhi.Description)

	return sayhi
}
