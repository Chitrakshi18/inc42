package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := map[string]string{"message": "Hello from Go API!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
	})

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		message := map[string]string{"message": "Hello World endpoint"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
	})

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
