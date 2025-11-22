package validators

import (
	"regexp"
	"retail_workflow/internal/shared/errors"
)

func ValidatePassword(password string) error {
	patterns := []regexp.Regexp{
		*regexp.MustCompile(`.*[a-z].*`),
		*regexp.MustCompile(`.*[A-Z].*`),
		*regexp.MustCompile(`.*\d.*`),
		*regexp.MustCompile(`^.{8,}$`),
		*regexp.MustCompile(`^.{8,48}$`),
	}

	errs := []error{
		errors.Err400.ErrLowercaseLettersIsMissingInPassword,
		errors.Err400.ErrUppercaseLettersIsMissingInPassword,
		errors.Err400.ErrNumbersIsMissingInPassword,
		errors.Err400.ErrPasswordTooShort,
		errors.Err400.ErrPasswordTooLong,
	}

	for i, p := range patterns {
		if !p.MatchString(password) {
			return errs[i]
		}
	}

	return nil
}
