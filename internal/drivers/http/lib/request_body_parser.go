package lib

import (
	"encoding/json"
	"io"
	"net/http"
	"retail_flow/internal/shared/exceptions"
)

func ParseRequestBody[T any](r *http.Request) (T, error) {
	var t T

	raw, err := io.ReadAll(r.Body)

	if err != nil {
		return t, exceptions.InternalExceptions.ErrFailureInParsingTheRequestBody
	}

	err = json.Unmarshal(raw, &t)

	if err != nil {
		return t, exceptions.InternalExceptions.ErrFailureInParsingTheRequestBody
	}

	return t, nil
}
