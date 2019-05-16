package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		n, err := client.Incr("hits").Result()
		if err != nil {
			log.Printf("error from redis: %v", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "Oopsie: %v\n", err)
			return
		}
		fmt.Fprintf(w, "Hits: %v\n", n)
		return
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
