package lib

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"retail_flow/internal/shared/exceptions"
	"testing"

	"github.com/stretchr/testify/require"
)

type tester struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func TestParseRequestBody(t *testing.T) {
	t.Run("when parsed", func(t *testing.T) {
		body := bytes.NewBufferString(`{
			"firstName": "John",
			"lastName": "Doe"
		}`)

		req := httptest.NewRequest(http.MethodGet, "/", body)
		payload, err := ParseRequestBody[tester](req)

		require.NoError(t, err)
		require.IsType(t, tester{}, payload)
	})

	t.Run("when body is nil", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		payload, err := ParseRequestBody[tester](req)

		require.ErrorIs(
			t,
			exceptions.InternalExceptions.ErrFailureInParsingTheRequestBody,
			err,
		)

		require.Empty(t, payload)
	})
}
