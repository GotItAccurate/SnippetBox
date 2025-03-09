package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
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

	// Parse MySQL DSN string :
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	// Parsing the flag :
	flag.Parse()

	// Create new loggers :
	// INFO logger :
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// ERROR logger :
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open DB :
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

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
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
