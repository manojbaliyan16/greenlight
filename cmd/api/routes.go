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
	//Now with the help of http.handler customize the error handler
	//1. Convert the notFoundResponse() helper to a http.Handler using the http.HandlerFunc() adapter and then set it into custom error handler for 404 notFoundResponse
	router.NotFound = http.HandlerFunc(app.notFoundResposne)
	// 2. likewise convert the methodNotallowedRespose()

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	//Register the methods using HandleFunc
	fmt.Println("Server is waiting Provided the routes ")
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
