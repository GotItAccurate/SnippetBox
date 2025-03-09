package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Sturct for sharing the application context :
type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {
	// Parse command line arguments :
	// Parse the PORT address :
	addr := flag.String("addr", ":4000", "HTTP PORT address")

	// Parsing the flag :
	flag.Parse()

	// Create new loggers :
	// INFO logger :
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// ERROR logger :
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Define application context
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

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

	// Define parameters for the HTTP server :
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// Starting the server :
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
