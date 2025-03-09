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

	// Define parameters for the HTTP server :
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	// Starting the server :
	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
