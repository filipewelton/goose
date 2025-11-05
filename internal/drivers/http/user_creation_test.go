package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	sharedLib "retail_flow/internal/shared/lib"
	"retail_flow/tests/generators"
	"testing"

	"github.com/stretchr/testify/require"
	"syreclabs.com/go/faker"
)

func TestUserCreation(t *testing.T) {
	var guardian sharedLib.APIGuardian
	var hash = faker.RandomString(64)

	sharedLib.SetupEnvironmentVariables()

	t.Run("when the user is successfully created", func(t *testing.T) {
		cardNumber := faker.Number().Number(7)
		name := faker.Name().Name()
		password := generators.GeneratePassword()

		body := bytes.NewBufferString(fmt.Sprintf(`{
			"cardNumber": "%s",
			"name": "%s",
			"password": "%s"
		}`, cardNumber, name, password))

		ts := httptest.NewServer(SetupServer())
		url := ts.URL + "/users"

		defer ts.Close()

		req, err := http.NewRequest(http.MethodPost, url, body)

		require.NoError(t, err)
		guardian.GenerateToken(hash)

		apiToken := guardian.GetToken()

		req.Header.Add("Authorization", fmt.Sprint("Bearer ", apiToken))

		client := http.Client{}
		res, err := client.Do(req)

		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, res.StatusCode)

		defer res.Body.Close()

		var responseBody map[string]any

		data, err := io.ReadAll(res.Body)

		require.NoError(t, err)

		err = json.Unmarshal(data, &responseBody)

		require.NoError(t, err)

		require.NotEmpty(t, responseBody["user"])
	})
}
