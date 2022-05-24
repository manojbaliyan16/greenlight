package main

import (
	"net/http"
)

// Declare a func which will write a plain text response with information about the a
// application status, env and version

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Println(" health check Route Provided  --7")
	/*fmt.Fprintln(w, "status:available")
	fmt.Fprintf(w, "enviorment: %s\n\n", app.config.env)
	fmt.Fprintf(w, "version : %s\n\n", version)*/
	/*js := `{"status": "available", "environment":%q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))*/
	// now we want to send the data in the form of JSON response so for that we have to make a map having data
	/*	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}*/
	/*// we will pass this map to JSON marshal and we will get back as the []byte containing the encoded JSON
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		return
	}
	js = append(js, '\n')
	// set Http header for a successfull response
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	// As we are going to send ;lots of JSON response so we can write our own WriteJSON() method into helper method
	*/
	// instead of above line we can call the fucntion defined into helper
	env := envelop{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(
			w,
			"The server encountered a problem and could not send process your request",
			http.StatusInternalServerError,
		)
	}
}
