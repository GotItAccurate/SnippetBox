package main

import (
	"log"
	"net/http"
)

func main() {
	// Create a new router/mux :
	mux := http.NewServeMux()

	// Create a file server for the static files :
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Route handler for the static files :
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Route handlers for the rest of the routes :
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/snippet/view", SnippetView)
	mux.HandleFunc("/snippet/create", SnippetCreate)

	// Starting the server :
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
