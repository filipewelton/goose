package apiguardian

import (
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var subject = environment.GetEnv(environment.API_GUARDIAN_SUBJECT)
var secret = environment.GetEnv(environment.API_GUARDIAN_SECRET)
var algorithm = jwt.SigningMethodHS256.Name

func Validate(token string) (jwt.Claims, error) {
	t, err := jwt.Parse(
		token,
		handleTokenValidation,
		jwt.WithSubject(subject),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{algorithm}),
	)

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    401,
			Context: "API toke validation",
			Reason:  err.Error(),
		})

		return nil, errors.Error401.ErrUnauthenticatedAPIAccess
	}

	return t.Claims, nil
}

func handleTokenValidation(t *jwt.Token) (any, error) {
	return []byte(secret), nil
}

func Generate() (string, error) {
	now := time.Now()

	claims := &jwt.RegisteredClaims{
		Subject:   subject,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 24 * 7)),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString([]byte(secret))

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "API token generation",
			Reason:  err.Error(),
		})

		return "", errors.Error500.ErrAPITokenGenerationFailure
	}

	return token, nil
}
