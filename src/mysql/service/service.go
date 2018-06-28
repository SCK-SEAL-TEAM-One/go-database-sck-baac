package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
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
	defer db.Close()

	if r.Method == http.MethodPost {
		description = r.FormValue("description")
		insertForm, err := db.Prepare("INSERT INTO sayhi(description) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(description)
	}
	w.Write([]byte("Create " + description + " success\n"))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()

	if r.Method == http.MethodDelete {
		id := r.FormValue("id")
		deleteForm, err := db.Prepare("DELETE FROM sayhi WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		w.Write([]byte("Deleted " + id + " success\n"))
		return
	}
	w.Write([]byte("fail"))
}
