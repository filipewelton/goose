package dto

import (
	"regexp"
	"retail_flow/internal/shared/exceptions"
)

type UserCreationDTO struct {
	CardNumber string `json:"cardNumber"`
	Name       string `json:"name"`
	Password   string `json:"password"`
}

func (u *UserCreationDTO) ValidateCardNumber() error {
	re := regexp.MustCompile(`^\d{6,7}$`)

	if !re.MatchString(u.CardNumber) {
		return exceptions.BadRequestExceptions.ErrInvalidCardNumber
	}

	return nil
}

func (u *UserCreationDTO) ValidateName() error {
	re := regexp.MustCompile(`^[a-zA-Z\s.]{2,50}$`)

	if !re.MatchString(u.Name) {
		return exceptions.BadRequestExceptions.ErrInvalidName
	}

	return nil
}

func (u *UserCreationDTO) ValidatePassword() error {
	lowercaseLetters := regexp.MustCompile(`.*[a-z].*`)
	uppercaseLetters := regexp.MustCompile(`.*[A-Z].*`)
	digits := regexp.MustCompile(`.*\d.*`)
	specialCharacters := regexp.MustCompile(`.*[!@#\$%\^&\*].*`)

	if len(u.Password) < 8 || len(u.Password) > 48 {
		return exceptions.BadRequestExceptions.ErrPasswordInvalidLength
	} else if !lowercaseLetters.MatchString(u.Password) {
		return exceptions.BadRequestExceptions.ErrPasswordRequiresLowercaseLetter
	} else if !uppercaseLetters.MatchString(u.Password) {
		return exceptions.BadRequestExceptions.ErrPasswordRequiresUppercaseLetter
	} else if !digits.MatchString(u.Password) {
		return exceptions.BadRequestExceptions.ErrPasswordRequiresDigit
	} else if !specialCharacters.MatchString(u.Password) {
		return exceptions.BadRequestExceptions.ErrPasswordRequiresSpecialCharacter
	}

	return nil
}
