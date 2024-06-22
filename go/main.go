package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Welcome from Go API!")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, _ *http.Request) {
		message := map[string]string{"message": "This is hello endpoint from Go API!"}
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(message)
		if err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			return
		}
	})

	port := ":8080"
	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
