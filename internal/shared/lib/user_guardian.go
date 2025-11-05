package lib

import (
	"os"
	"retail_flow/internal/shared/exceptions"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type UserGuardian struct {
	secret       []byte
	accessToken  string
	refreshToken string
}

func (u *UserGuardian) GenerateAccessToken(
	userID string,
	scopes []string,
) error {
	token, err := u.handleTokeGeneration(
		userID,
		scopes,
		time.Now().Add(time.Hour*24*7).Unix(),
	)

	if err != nil {
		return err
	}

	u.accessToken = token
	return nil
}

func (u *UserGuardian) GenerateRefreshToken(
	userID string,
	scopes []string,
) error {
	token, err := u.handleTokeGeneration(
		userID,
		scopes,
		time.Now().Add(time.Hour*24).Unix(),
	)

	if err != nil {
		return err
	}

	u.refreshToken = token
	return nil
}

func (u *UserGuardian) handleTokeGeneration(
	userID string,
	scopes []string,
	expiration int64,
) (string, error) {
	var log = Logger()

	defer log.Sync()

	if err := u.getSecret(); err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":    expiration,
		"sub":    userID,
		"scopes": scopes,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := accessToken.SignedString(u.secret)

	if err != nil {
		log.Error(
			exceptions.InternalExceptions.ErrJWTGenerationFailed.Error(),
			zap.Error(err),
		)

		return "", exceptions.InternalExceptions.ErrJWTGenerationFailed
	}

	return str, nil
}

func (u *UserGuardian) getSecret() error {
	if len(u.secret) != 0 {
		return nil
	}

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		return exceptions.InternalExceptions.ErrEnvironmentVariableIsMissing
	}

	u.secret = []byte(secret)
	return nil
}

func (u *UserGuardian) GetAccessToken() string {
	return u.accessToken
}

func (u *UserGuardian) GetRefreshToken() string {
	return u.refreshToken
}
