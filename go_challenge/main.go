package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Hepsiburada from Emre</h1>")
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server is starting...")
	http.ListenAndServe(":11130", nil)
}
