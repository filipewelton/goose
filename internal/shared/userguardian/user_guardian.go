package userguardian

import (
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = environment.GetEnv(environment.USER_GUARDIAN_SECRET)
var algorithm = jwt.SigningMethodHS256.Name

func Validate(token string) (jwt.Claims, error) {
	t, err := jwt.Parse(
		token,
		handleTokenValidation,
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{algorithm}),
	)

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    401,
			Context: "User token validation",
			Reason:  err.Error(),
		})

		return nil, errors.Error401.ErrInvalidUserCredentials
	}

	return t.Claims, nil
}

func handleTokenValidation(t *jwt.Token) (any, error) {
	return []byte(secret), nil
}

func GenerateAccessToken(userId string) (string, error) {
	return handleTokenGeneration(
		"User access token generation",
		userId,
		time.Hour*24*7,
	)
}

func GenerateRefreshToken(userId string) (string, error) {
	return handleTokenGeneration(
		"User refresh token generation",
		userId,
		time.Hour*24,
	)
}

func handleTokenGeneration(
	context string,
	userId string,
	exp time.Duration,
) (string, error) {
	now := time.Now()

	claims := &jwt.RegisteredClaims{
		Subject:   userId,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(exp)),
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secret))

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: context,
			Reason:  err.Error(),
		})

		return "", errors.Error500.ErrUserCredentialGenerationFailure
	}

	return token, nil
}
