package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Create a new router/mux :
	mux := http.NewServeMux()

	// Create a file server for the static files :
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Route handler for the static files :
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Route handlers for the rest of the routes :
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.SnippetView)
	mux.HandleFunc("/snippet/create", app.SnippetCreate)

	return mux
}
