package main

import (
	"net/http"
	"postgresql/api"

	_ "github.com/lib/pq"
)

func main() {
	http.HandleFunc("/read", api.Read())
	http.HandleFunc("/readById", api.ReadById())
	http.HandleFunc("/add", api.Read())
	http.HandleFunc("/edit", api.Read())
	http.HandleFunc("/delete", api.Read())
	http.ListenAndServe(":8080", nil)
}
