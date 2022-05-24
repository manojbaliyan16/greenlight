package main

import (
	"fmt"
	"net/http"
	"time"

	"greenloght.manoj.dev/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Movie route provided ")
	fmt.Fprintf(w, "create a new Movie of your choice ")

}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	// when httprouter parsing a request any interpolated URL parameters will be
	// stored in the request context
	// For That we can use ParamsFromContext() function to retrieve a slice containing these parameters names and values
	fmt.Println("Show Movie Handler provided ")
	fmt.Println("Is request reaching¸ up to here ")
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
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
		app.logger.Println(err)
		http.Error(
			w,
			"The server encountered a problem and could not proccess your request",
			http.StatusInternalServerError,
		)

	}

}
