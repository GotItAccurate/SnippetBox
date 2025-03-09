package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Parsing command line arguments :
	// Parsing the PORT address :
	addr := flag.String("addr", ":4000", "HTTP PORT address")

	// Parsing the flag :
	flag.Parse()

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
	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
