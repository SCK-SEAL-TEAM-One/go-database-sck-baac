package main

import (
	"log"
	"mysql/handler"
	"net/http"
)

func main() {
	api := handler.Api{
		DBDriver:   "mysql",
		DBUsername: "root",
		DBPassword: "sckshuhari",
		DBName:     "sckseal",
	}
	http.HandleFunc("/create", api.CreateMessageHandler)
	http.HandleFunc("/delete", api.Delete)
	http.HandleFunc("/update", api.Update)
	http.HandleFunc("/readbyid", api.ReadById)
	http.HandleFunc("/read", api.Read)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
