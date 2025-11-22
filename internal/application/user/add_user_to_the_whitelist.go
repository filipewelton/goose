package user

import (
	domain "retail_workflow/internal/domain/user"
	"retail_workflow/internal/shared/validators"
)

type WhitelistInclusion struct {
	WhitelistRepository domain.WhitelistRepository
	Payload             domain.WhitelistInclusionDTO
}

func AddUserToTheWhitelist(params WhitelistInclusion) error {
	var employeeId = params.Payload.EmployeeId

	if err := validators.ValidateEmployeeId(employeeId); err != nil {
		return err
	}

	repository := params.WhitelistRepository
	has, err := repository.Has(employeeId)

	if err == nil || has {
		return nil
	}

	return repository.Insert(employeeId)
}
