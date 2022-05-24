package main

import (
	"fmt"
	"net/http"
)

//logError is a generic helper for logging an error message.
func (app *application) logError(
	r *http.Request,
	err error,
) {
	app.logger.Println(err)
}

//errorResponse is a generic helper for sending JSON formatted error
func (app *application) errorResponse(
	w http.ResponseWriter,
	r *http.Request,
	status int,
	message interface{},
) {
	env := envelop{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)

	}
}

//ServerErrorrespose will be used when our application encounters an unexpected Problem at run time. It log the details error
// message Then uses the errorResposne() helper to send a 500 iternal Server Error

func (app *application) serverErrorResponse(
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

//notfoundResposne method will be used to send a 404 not found status
//status code and JSON response to the client

func (app *application) notFoundResposne(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := "The requested source  could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

//methodNotAllowedResponse will be used to send a 405 Method Not allowed
//Status code and JSON response to the client
func (app *application) methodNotAllowedResponse(
	w http.ResponseWriter,
	r *http.Request,
) {
	message := fmt.Sprintf("the %s method is not supposrted for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}
