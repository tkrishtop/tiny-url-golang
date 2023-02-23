package main

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var longToShort = make(map[string]string)
var shortToLong = make(map[string]string)

const baseAPIURL = "https://tiny.com/"
const shortCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/url/v1/long2short", LongToShortHandler).Methods("POST")
	r.HandleFunc("/url/v1/short2long/{shortURL}", ShortToLongHandler).Methods("GET")
	http.ListenAndServe(":3003", r)
}

func LongToShortHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	longURL := string(body)
	fmt.Println("Request to encode long URL:", longURL)
	shortURL := Encode(longURL)
	fmt.Println("Encoded:", shortURL)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte(shortURL))
}

func ShortToLongHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	fmt.Println("Request to decode short URL:", shortURL)
	longURL := Decode(shortURL)
	fmt.Println("Encoded:", longURL)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(200)
	w.Write([]byte(longURL))
}

func Encode(longURL string) string {
	hashURL, ok := longToShort[longURL]

	if ok { // if this URL is already in the hashmap
		return baseAPIURL + hashURL
	} else { // if not, encode it
		h := fnv.New32a()
		h.Write([]byte(longURL))
		hashValue := h.Sum32()

		// Convert the hash value to base 62
		var sb strings.Builder
		base := uint32(len(shortCharset))
		for hashValue > 0 {
			sb.WriteByte(shortCharset[hashValue%base])
			hashValue /= base
		}
		hashURL = sb.String()
		longToShort[longURL] = hashURL
		shortToLong[hashURL] = longURL

		return baseAPIURL + hashURL
	}
}

func Decode(shortURL string) string {
	longURL, ok := shortToLong[shortURL]

	if ok { // if this short URL is in the hashmap
		return longURL
	} else { // if not, return not_found
		return "not found"
	}
}
