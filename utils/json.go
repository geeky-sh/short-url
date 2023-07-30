package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func JSNDecode(r *http.Request, d interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&d)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return NewAppErr(msg, ERR_JSON_PARSE)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintln("Request body contains badly-formed JSON")
			return NewAppErr(msg, ERR_JSON_PARSE)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return NewAppErr(msg, ERR_JSON_PARSE)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return NewAppErr(msg, ERR_JSON_PARSE)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return NewAppErr(msg, ERR_JSON_PARSE)

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return NewAppErr(msg, ERR_JSON_PARSE)

		default:
			return NewAppErr("Unknown Error", ERR_JSON_PARSE)
		}
	}
	return nil
}
