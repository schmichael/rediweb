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

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		status, err := client.Ping().Result()
		if err != nil {
			log.Printf("ping: error from redis: %v", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "Ping oopsie: %v\n", err)
			return
		}
		fmt.Fprintf(w, "Ok\nStatus: %v\n", status)
		return
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		n, err := client.Incr("hits").Result()
		if err != nil {
			log.Printf("error from redis: %v", err)
			w.WriteHeader(500)
			fmt.Fprintf(w, "oopsie: %v\n", err)
			return
		}
		fmt.Fprintf(w, "Hits: %v\n", n)
		return
	})

	addr := "127.0.0.1:8080"
	log.Printf("Listening on http://%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
