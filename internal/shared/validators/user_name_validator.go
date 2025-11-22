package validators

import (
	"regexp"
	"retail_workflow/internal/shared/errors"
)

func ValidateUserName(name string) error {
	pattern := regexp.MustCompile(`^[a-zA-Z.\s]+$`)

	if !pattern.MatchString(name) {
		return errors.Err400.ErrInvalidUserName
	}

	return nil
}
