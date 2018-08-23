package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")
		if len(key) == 0 {
			w.WriteHeader(http.StatusOK)
			if _, err := w.Write([]byte("missing 'key' parameter")); err != nil {
				log.Println("error writing response: ", err)
			}
			return
		}

		if r.Method == http.MethodPost {
			if err := client.Set(key, r.FormValue("value"), 0).Err(); err != nil {
				log.Println("Redis set error: ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			return
		}

		value, err := client.Get(key).Result()
		if err != nil {
			log.Println("Redis get error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte(value)); err != nil {
			log.Println("error writing response: ", err)
		}
	}))
}
