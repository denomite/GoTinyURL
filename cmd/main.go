package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	// interal code !

	r := mux.NewRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
