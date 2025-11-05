package valueobjects

import (
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	value string
}

func (e *Password) Hash(plaintext string) error {
	var log = lib.Logger()

	defer log.Sync()

	hash, err :=
		bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)

	if err != nil {
		log.Error(
			exceptions.InternalExceptions.ErrPasswordHashingFailed.Error(),
			zap.Error(err),
		)

		return exceptions.InternalExceptions.ErrPasswordHashingFailed
	}

	e.value = string(hash)
	return nil
}

func (e *Password) GetValue() string {
	return e.value
}

func (e *Password) SetValue(value string) {
	e.value = value
}
