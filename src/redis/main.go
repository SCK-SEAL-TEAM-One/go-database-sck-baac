package main

import (
	"net/http"
	"redis/api"
	"redis/crud"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	rc := crud.RedisClient{Client: client}
	http.HandleFunc("/helloworld/readAll", api.GetReadAll(rc.ReadAllKey))
	http.HandleFunc("/helloworld/readById", api.GetReadById(rc.ReadById))
	http.HandleFunc("/redis/read", api.GetRead(rc.ReadKey))

	http.HandleFunc("/redis/create", api.GetBreadCreate(rc.WriteKey))

	http.HandleFunc("/redis/delete", api.GetDelete(rc.DeleteKey))

	http.ListenAndServe(":3000", nil)

}
