package main

import (
	"net/http"
)

// Declare a func which will write a plain text response with information about the a
// application status, env and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// Now introducing the error handling and sending JSOn response to the client as per the error
	env := envelop{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		// use the new serverErrorResponse() helper
		app.serverErrorResponse(w, r, err)
	}
}
