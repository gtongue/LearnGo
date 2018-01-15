package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the index function")
}
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about function")
}
