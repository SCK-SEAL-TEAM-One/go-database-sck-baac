package crud

import (
	"database/sql"
	"fmt"
	"postgresql/model"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "sckseal"
)

func ConnectDB() (db *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func ReadSayhi(id string, postgresURL string) model.Sayhi {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM sayhi")
	var listsayhi [10]model.Sayhi
	for rows.Next() {
		var id int
		var description string
		err = rows.Scan(&id, &description)
	}
	return listsayhi
}

// func ReadSayhiByID() string {
// 	return "ggggg"
// }

// func ReadSayhi() string {
// 	return "ggggg"
// }
