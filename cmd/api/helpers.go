package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func (app *application) readJSON(
	w http.ResponseWriter,
	r *http.Request,
	dst interface{},
) error {
	//Decode the request body into the target destination
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// if there is an error during the decoding start the triage
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {
		//erros.As --> whether the error has the type

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly formateed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formatted JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type at character %d", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err

		}

	}
	return nil
}
