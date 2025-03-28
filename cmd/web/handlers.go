package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GotItAccurate/SnippetBox/internal/models"
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

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Output for the specified ID :
	fmt.Fprintf(w, "%+v", snippet)
}

// Handler for POST : '/snippet/create'
func (app *application) SnippetCreate(w http.ResponseWriter, r *http.Request) {
	// Check for POST method :
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
