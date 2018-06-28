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
	db := dbConnection()
	defer db.Close()

	if r.Method == http.MethodPost {
		description := r.FormValue("description")
		insertForm, err := db.Prepare("INSERT INTO sayhi(description) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(description)
		w.Write([]byte("Created " + description + " success\n"))
		return
	}
	w.Write([]byte("fail"))

}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()

	if r.Method == http.MethodPost {
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

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()
	if r.Method == http.MethodPost {
		description := r.FormValue("description")
		id := r.FormValue("id")
		insertForm, err := db.Prepare("UPDATE sayhi SET description=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(description, id)
		w.Write([]byte("Updated description :" + description + " id :" + id + " success\n"))
		return
	}
	w.Write([]byte("fail"))

}

func ReadById(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		deleteForm, err := db.Prepare("SELECT description FROM sayhi WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		rows, _ := deleteForm.Query(id)
		defer rows.Close()
		for rows.Next() {
			var description string
			rows.Scan(&description)
			w.Write([]byte(description))
		}
		return
	}
	w.Write([]byte("fail"))
}

func Read(w http.ResponseWriter, r *http.Request) {
	db := dbConnection()
	defer db.Close()
	if r.Method == http.MethodGet {
		deleteForm, err := db.Prepare("SELECT id, description FROM sayhi")
		if err != nil {
			panic(err.Error())
		}
		rows, _ := deleteForm.Query()
		defer rows.Close()
		for rows.Next() {
            var id string
			var description string
			rows.Scan(&id,&description)
			w.Write([]byte(id + ":" +description + "\n"))
		}
		return
	}
	w.Write([]byte("fail"))
}
