package adminuserauthentication

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/drivers/httpdriver/utils"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/userguardian"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const AccessTokenCookieName = "access_token"
const RefreshTokenCookieName = "refresh_token"

var userRepository user.UserRepository

func SetMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, refreshToken, err := getTokens(ctx)

		if err != nil {
			utils.Render(err, ctx)
			return
		}

		claim, err := validateTokens(accessToken, refreshToken)

		if err != nil {
			utils.Render(err, ctx)
			return
		}

		subject, err := claim.GetSubject()

		if err != nil {
			utils.Render(errors.Error401.ErrUserUnauthenticated, ctx)
			return
		}

		err = checkIfTheUserExists(subject)

		if err != nil {
			utils.Render(err, ctx)
			return
		}

		ctx.Next()
	}
}

func getTokens(ctx *gin.Context) (string, string, error) {
	accessToken, err1 := ctx.Cookie(AccessTokenCookieName)
	refreshToken, err2 := ctx.Cookie(RefreshTokenCookieName)

	if err1 != nil || err2 != nil {
		return "", "", errors.Error401.ErrUserUnauthenticated
	}

	return accessToken, refreshToken, nil
}

func validateTokens(accessToken, refreshToken string) (jwt.Claims, error) {
	claim, err1 := userguardian.Validate(accessToken)
	_, err2 := userguardian.Validate(refreshToken)

	if err1 != nil || err2 != nil {
		return nil, errors.Error401.ErrUserUnauthenticated
	}

	return claim, nil
}

func SetUserRepository(repository user.UserRepository) {
	userRepository = repository
}

func checkIfTheUserExists(subject string) error {
	if userRepository == nil {
		return errors.Error500.ErrUndefinedUserRepository
	}

	_, err := userRepository.FindById(subject)

	return err
}
