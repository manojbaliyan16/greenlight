package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Configure the server with endpoint /v1/healthcheck/

const version = "1.0.1"

//config to hold the configuration setting for out application
// For Now only configuration setting is the network port
type config struct {
	port int
	env  string
}

// application to hold the dependencies for our http handler, helpers and middleware
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// read the value of the port and enc command line flags config struct
	//fmt.Println("Configure the API server Addr--1")
	flag.IntVar(
		&cfg.port,
		"port",
		4000,
		"API server port",
	)
	//fmt.Println("Configure the Enviorment Variable --2")
	flag.StringVar(
		&cfg.env,
		"env",
		"dev",
		"Enviorment(development|staging|production)",
	)
	// logger to write the message to stdout
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	// instance of application struct having config struct and logger
	app := &application{
		config: cfg,
		logger: logger,
	}
	//New serverMux and add a /v1/healthcheck route which dispatches teh requests
	// to the healthcheckhandler
	/*fmt.Println("Create the Http serbver and provide a route --4")
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	*/
	// declare an HTTP server with some sensible timeout settings,
	//fmt.Println("Give the Time Out   --6")
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	// start the server
	//fmt.Println("Print the log --7")
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
