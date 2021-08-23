package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

func sources(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		`{"Name":"source 1","Body":"Hello","Time":1294706395881547000}`,
	))
}

func authentications(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(
		`{"Name":"auth 1","Body":"Hello","Time":1294706395881547000}`,
	))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/authentications", authentications)
	mux.HandleFunc("/sources", sources)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
