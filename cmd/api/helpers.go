package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type envelop map[string]interface{}

func (app *application) readIDParam(r *http.Request) (int64, error) {
	fmt.Println("To read the param or parse the URL we came ReadIDParam")
	param := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(param.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("Invalid id param")
	}
	return id, nil

	// writeJSON() helper method to send the response that will take the destination
	// http.ResposneWriter, the http status code to send the data to encode to JSON and a
	// header map containing any additional information HTTP headers we want to include in the resposne

}
func (app *application) writeJSON(
	w http.ResponseWriter,
	status int,
	data envelop,
	headers http.Header,
) error {
	// Encode data to JSON
	js, err := json.MarshalIndent(data, "", "\t")
	//js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// append a new line
	js = append(js, '\n')
	// add any header we want to include
	// through the header map and add each header to the
	// http.ResposneWriter header map
	for key, value := range headers {
		w.Header()[key] = value
	}
	// add the content type header and then write the status code
	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(status)
	w.Write(js)
	//
	return nil

}
