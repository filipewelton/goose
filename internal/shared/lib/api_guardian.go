package lib

import (
	"os"
	"retail_flow/internal/shared/exceptions"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type APIGuardian struct {
	secret []byte
	token  string
}

func (a *APIGuardian) GenerateToken(frontendHash string) error {
	var log = Logger()

	defer log.Sync()

	if err := a.getSecret(); err != nil {
		return err
	}

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": frontendHash,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := accessToken.SignedString(a.secret)

	if err != nil {
		log.Error(
			exceptions.InternalExceptions.ErrJWTGenerationFailed.Error(),
			zap.Error(err),
		)

		return exceptions.InternalExceptions.ErrJWTGenerationFailed
	}

	a.token = str
	return nil
}

func (a *APIGuardian) getSecret() error {
	if len(a.secret) != 0 {
		return nil
	}

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return exceptions.InternalExceptions.ErrEnvironmentVariableIsMissing
	}

	a.secret = []byte(secret)
	return nil
}

func (a *APIGuardian) GetToken() string {
	return a.token
}

func (a *APIGuardian) Validate(token string) error {
	var log = Logger()

	defer log.Sync()

	_, err := jwt.ParseWithClaims(
		token,
		jwt.MapClaims{},
		func(t *jwt.Token) (any, error) {
			exp, err := t.Claims.GetExpirationTime()

			if err != nil {
				return nil, err
			} else if time.Now().After(exp.Time) {
				return nil, exceptions.UnauthorizedExceptions.ErrAPITokenExpired
			} else if err = a.getSecret(); err != nil {
				return nil, err
			}

			return a.secret, nil
		},
	)

	if err != nil {
		log.Error(
			exceptions.UnauthenticatedExceptions.ErrInvalidAPIToken.Error(),
			zap.Error(err),
		)

		return exceptions.UnauthenticatedExceptions.ErrInvalidAPIToken
	}

	return nil
}
