package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey</h1>")
	fmt.Fprintf(w, "<h2>World</h2>")
	fmt.Fprintf(w, "<h3>Wassup</h3>")
	fmt.Fprintf(w, "<p>This %s is using %s</p>", "program", "variables")
}

func main() {
	http.HandleFunc("/", index_handler)
	fmt.Println("Server Started at port 8000")
	http.ListenAndServe(":8000", nil)
}
