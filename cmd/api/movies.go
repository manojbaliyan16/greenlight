package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"greenloght.manoj.dev/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Movie route provided ")
	fmt.Fprintf(w, "create a new Movie of your choice ")
	//JSON Decoding : A client request to the end point by sending the JSON request how can we decode it
	// Below info supposed to be in HTTP request body
	var input struct {
		Title   string
		Year    int32
		Runtime int32
		Genere  []string
	}
	// initialize a new json.Decoder instance whoch reads from the request body
	// then decode to get body content into input struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// dump the inout content of the struct into HTTP resposne
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// when httprouter parsing a request any interpolated URL parameters will be
	// stored in the request context
	// For That we can use ParamsFromContext() function to retrieve a slice containing these parameters names and values
	fmt.Println("Show Movie Handler provided ")
	fmt.Println("Is request reaching¸ up to here ")
	id, err := app.readIDParam(r)
	if err != nil {
		//Use the new NotFoundResponse() helper
		app.notFoundResposne(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of the movie %d\n", id)
	// create an instance of movie data

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}
	//endcode the struct to json and send it as HTTP response
	err = app.writeJSON(w, http.StatusOK, envelop{"movie": movie}, nil)
	if err != nil {
		//Use the new serverErrorResponse() helper
		app.serverErrorResponse(w, r, err)
	}

}
