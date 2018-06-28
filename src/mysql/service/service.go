package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    "net/http"
    "fmt"
)

func dbConnection() (db *sql.DB) {
    // Hard Code Database Config
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "sckshuhari"
    dbName := "sckseal"
    sqlCommand := fmt.Sprintf("%s:%s@/%s", dbUser, dbPass, dbName)
	db, err := sql.Open(dbDriver, sqlCommand)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
    var description string
    db := dbConnection()
    
	if r.Method == "POST" {
		description = r.FormValue("description")
		insertForm, err := db.Prepare("INSERT INTO sayhi(description) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(description)
	}
    defer db.Close()
    w.Write([]byte("Create " + description + " success\n"))
}


