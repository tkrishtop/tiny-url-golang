package main

import (
	"fmt"
	"net/http"
	"strings"
)

var longToShort = make(map[string]string)
var shortToLong = make(map[string]string)

func encode(longUrl string) string {
	shortUrl, ok := longToShort[longUrl]
	var response string
	if ok {
		response = shortUrl
	} else {
		response = "not_found"
	}
	return response
}

func home(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	longUrls := query["long"]
	encodedUrls := make([]string, len(longUrls))
	for idx, longUrl := range longUrls {
		fmt.Println(longUrl)
		encodedUrls[idx] = encode(longUrl)
	}

	// shortUrls := query["short"]

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	response := "Encoded long URLs: " + strings.Join(encodedUrls, ", ") + "\n"
	w.Write([]byte(response))
}

func handleRequests() {
	http.HandleFunc("/url", home)
	http.ListenAndServe(":3003", nil)
}

func main() {
	handleRequests()
}
