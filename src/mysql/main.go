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
	http.HandleFunc("/helloworld/add", api.CreateMessageHandler)
	http.HandleFunc("/helloworld/delete", api.Delete)
	http.HandleFunc("/helloworld/edit", api.Update)
	http.HandleFunc("/helloworld/readById", api.ReadById)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
