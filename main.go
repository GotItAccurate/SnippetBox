package main

import (
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from SnippetBox."))
}

func SnippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetView."))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from SnippetCreate."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
