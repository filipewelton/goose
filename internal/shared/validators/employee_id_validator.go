package validators

import (
	"regexp"
	"retail_workflow/internal/shared/errors"
)

func ValidateEmployeeId(employeeId string) error {
	pattern := regexp.MustCompile(`^\d{6,7}$`)

	if !pattern.MatchString(employeeId) {
		return errors.Err400.ErrInvalidEmployeeId
	}

	return nil
}
