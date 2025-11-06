package user

import (
	"retail_flow/internal/application/dto"
	"retail_flow/internal/application/interfaces"
	"retail_flow/internal/application/mappers"
	"retail_flow/internal/domain/constants"
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"
	"retail_flow/internal/shared/typings"
)

type UserCreationParams struct {
	UserRepository      interfaces.UserRepository
	WhitelistRepository interfaces.WhitelistRepository
	Payload             dto.UserCreationDTO
	entity              entities.UserEntity
	accessToken         string
	refreshToken        string
}

type UserCreationResult struct {
	User         typings.UserEntityMapped `json:"user"`
	AccessToken  string                   `json:"-"`
	RefreshToken string                   `json:"-"`
}

func Create(params UserCreationParams) (UserCreationResult, error) {
	operations := []func() error{
		params.checkIfTheCardNumberIsOnTheWhitelist,
		params.handleValidations,
		params.checkIfTheCardNumberIsAlreadyRegistered,
		params.createEntity,
		params.persist,
		params.generateTokens,
	}

	for _, operation := range operations {
		if err := operation(); err != nil {
			return UserCreationResult{}, err
		}
	}

	return UserCreationResult{
		User:         mappers.MapUserEntity(params.entity),
		AccessToken:  params.accessToken,
		RefreshToken: params.refreshToken,
	}, nil
}

func (p *UserCreationParams) checkIfTheCardNumberIsOnTheWhitelist() error {
	cardNumber := p.Payload.CardNumber
	_, err := p.WhitelistRepository.Has(cardNumber)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserCreationParams) handleValidations() error {
	validations := []func() error{
		p.Payload.ValidateCardNumber,
		p.Payload.ValidateName,
		p.Payload.ValidatePassword,
	}

	for _, validation := range validations {
		if err := validation(); err != nil {
			return err
		}
	}

	return nil
}

func (p *UserCreationParams) checkIfTheCardNumberIsAlreadyRegistered() error {
	cardNumber := p.Payload.CardNumber
	result, _ := p.UserRepository.FindByCardNumber(cardNumber)

	if result != (entities.UserEntity{}) {
		return exceptions.ConflictExceptions.ErrCardNumberAlreadyRegistered
	}

	return nil
}

func (p *UserCreationParams) createEntity() error {
	var entity entities.UserEntity

	if err := entity.New(p.Payload); err != nil {
		return err
	}

	p.entity = entity
	return nil
}

func (p *UserCreationParams) persist() error {
	if err := p.UserRepository.Insert(p.entity); err != nil {
		return err
	}

	return nil
}

func (u *UserCreationParams) generateTokens() error {
	guardian := lib.UserGuardian{}
	scopes := []string{constants.CALCULATOR, constants.EXPIRATION_DATE}

	err := guardian.GenerateAccessToken(
		u.entity.ID.GetValue(),
		scopes,
	)

	if err != nil {
		return err
	}

	err = guardian.GenerateRefreshToken(
		u.entity.ID.GetValue(),
		scopes,
	)

	if err != nil {
		return err
	}

	u.accessToken = guardian.GetAccessToken()
	u.refreshToken = guardian.GetRefreshToken()

	return nil
}
