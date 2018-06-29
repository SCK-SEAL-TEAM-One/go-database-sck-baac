package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Api struct {
	DBDriver   string
	DBUsername string
	DBPassword string
	DBName     string
}

type Sayhi struct {
	Id          string `json:"Id"`
	Description string `json:Description`
}

func MakeJson(id, description string) Sayhi {
	return Sayhi{
		Id:          id,
		Description: description,
	}
}

func (a Api) dbConnection() (db *sql.DB) {
	sqlCommand := fmt.Sprintf("%s:%s@/%s", a.DBUsername, a.DBPassword, a.DBName)
	db, err := sql.Open(a.DBDriver, sqlCommand)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func (a Api) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
	defer db.Close()

	if r.Method == http.MethodPost {
		description := r.FormValue("description")
		insertForm, err := db.Prepare("INSERT INTO sayhi(description) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insertForm.Exec(description)
		getLastID, err := db.Prepare("SELECT LAST_INSERT_ID()")
		rows, _ := getLastID.Query()
		defer rows.Close()
		var lastID uint64
		rows.Next()
		rows.Scan(&lastID)

		return
	}
	w.Write([]byte("fail"))

}

func (a Api) Delete(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
	defer db.Close()

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		deleteForm, err := db.Prepare("DELETE FROM sayhi WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		deleteForm.Exec(id)
		w.Write([]byte("Deleted success\n"))
		return
	}
	w.Write([]byte("fail"))
}

func (a Api) Update(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
	defer db.Close()
	if r.Method == http.MethodPost {
		description := r.FormValue("description")
		id := r.FormValue("id")
		insertForm, err := db.Prepare("UPDATE sayhi SET description=? WHERE id=?")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		insertForm.Exec(description, id)
		createResponse, _ := json.Marshal(MakeJson(id, description))
		w.Write(createResponse)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (a Api) ReadById(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
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

func (a Api) Read(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
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
			rows.Scan(&id, &description)
			w.Write([]byte(id + ":" + description + "\n"))
		}
		return
	}
	w.Write([]byte("fail"))
}
