package repositories

import "retail_flow/internal/shared/exceptions"

type InMemoryWhitelistRepository struct{}

var whitelistDB = map[string]string{}

func (InMemoryWhitelistRepository) Insert(cardNumber string) error {
	whitelistDB[cardNumber] = "OK"
	return nil
}

func (InMemoryWhitelistRepository) Has(cardNumber string) (bool, error) {
	if whitelistDB[cardNumber] == "" {
		return false, exceptions.NotFoundExceptions.ErrPreRegisterNotFound
	}

	return true, nil
}

func (InMemoryWhitelistRepository) Delete(cardNumber string) error {
	delete(whitelistDB, cardNumber)
	return nil
}
