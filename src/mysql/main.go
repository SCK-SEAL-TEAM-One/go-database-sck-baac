package main

import (
	"log"
	"mysql/service"
	"net/http"
)

func main() {
	http.HandleFunc("/create", service.CreateMessageHandler)
	http.HandleFunc("/delete", service.Delete)
	http.HandleFunc("/update", service.Update)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
