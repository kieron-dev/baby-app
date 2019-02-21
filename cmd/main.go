package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s\n", r.UserAgent())
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
