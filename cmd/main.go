package main

import (
	"context"
	"gotinyurl/internal/service"
	"gotinyurl/internal/storage"
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

	store := storage.NewRedisStore(rdb, ctx)
	service := service.NewService(store)

	r := mux.NewRouter()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
