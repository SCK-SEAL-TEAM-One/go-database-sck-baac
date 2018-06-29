package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

func (a Api) dbConnection() *sql.DB {
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
			w.WriteHeader(http.StatusInternalServerError)
		}
		insertForm.Exec(description)
		getLastID, err := db.Prepare("SELECT LAST_INSERT_ID()")
		rows, _ := getLastID.Query()
		defer rows.Close()
		var lastID uint64
		rows.Next()
		rows.Scan(&lastID)
		CreateResponse, _ := json.Marshal(MakeJson(strconv.FormatUint(lastID, 10), description))
		w.Write(CreateResponse)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (a Api) Delete(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
	defer db.Close()

	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		deleteForm, err := db.Prepare("DELETE FROM sayhi WHERE id=?")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		deleteForm.Exec(id)
		w.Write([]byte("Deleted success\n"))
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
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
		UpdateResponse, _ := json.Marshal(MakeJson(id, description))
		w.Write(UpdateResponse)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func (a Api) ReadById(w http.ResponseWriter, r *http.Request) {
	db := a.dbConnection()
	defer db.Close()
	if r.Method == http.MethodPost {
		id := r.FormValue("id")
		deleteForm, err := db.Prepare("SELECT id, description FROM sayhi WHERE id=?")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		rows, _ := deleteForm.Query(id)
		defer rows.Close()
		for rows.Next() {
			var id string
			var description string
			rows.Scan(&id, &description)
			readByIDResponse, _ := json.Marshal(MakeJson(id, description))
			w.Write(readByIDResponse)
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}
