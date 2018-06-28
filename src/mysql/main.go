package main

import (
	"log"
	"mysql/service"
	"net/http"
)

func main() {
	http.HandleFunc("/create", service.CreateMessageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
