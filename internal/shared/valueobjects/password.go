package valueobjects

import (
	"retail_workflow/internal/shared/errors"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func (p *Password) Hash(plaintext string) error {
	hash, err :=
		bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)

	if err != nil {
		return errors.Error500.ErrPasswordHashGenerationFailure
	}

	*p = Password(hash)
	return nil
}

func (p *Password) Set(hash string) {
	*p = Password(hash)
}

func (p *Password) Get() string {
	return string(*p)
}

func (p *Password) Compare(plaintext string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*p), []byte(plaintext))

	if err != nil {
		return errors.Error401.ErrInvalidUserCredential
	}

	return nil
}
