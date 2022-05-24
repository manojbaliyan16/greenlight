package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routing with httprouter and put all rounting base methods here

func (app *application) routes() *httprouter.Router {
	// initiate a new httprouter router instance
	router := httprouter.New()
	//Register the methods using HandleFunc
	fmt.Println("Server is waiting Provided the routes ")
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
