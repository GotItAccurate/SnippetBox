package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Handler for GET : '/' :
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// Check for absolute route :
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Template files :
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Parseing the template files :
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Writing the parsed template to the ResponseWriter :
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}
}

// Handler for GET : '/snippet/view?id=%d' :
func (app *application) SnippetView(w http.ResponseWriter, r *http.Request) {
	// Parseing the query string :
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Output for the specified ID :
	fmt.Fprintf(w, "Display a specific snippet ID %d.", id)
}

// Handler for POST : '/snippet/create'
func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check for POST method :
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello from SnippetCreate."))
}
