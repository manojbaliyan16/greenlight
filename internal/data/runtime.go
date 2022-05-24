package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// Generate the string containing the movie runtime in the required format

	jsonValue := fmt.Sprintf("%d  mins", r)
	// Valid JSON string needs to be surrounded in a double quote
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}
