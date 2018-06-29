package main

import "net/http"

func main() {
	http.HandleFunc("/read", api)
	http.ListenAndServe(":3000", nil)
}
func api(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("JSONConnectData"))
}
