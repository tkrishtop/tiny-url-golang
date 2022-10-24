package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Welcome!")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3003", nil)
}

func main() {
	handleRequests()
}
