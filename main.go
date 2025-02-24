package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", Handler)

	log.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)

}
