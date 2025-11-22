package mocks

type InMemoryWhitelistRepository struct{}

var inMemoryWhiteList = make(map[string]string)

func (InMemoryWhitelistRepository) Insert(employeeId string) error {
	inMemoryWhiteList[employeeId] = "OK"
	return nil
}

func (InMemoryWhitelistRepository) Has(employeeId string) (bool, error) {
	has := inMemoryWhiteList[employeeId] == "OK"
	return has, nil
}

func (InMemoryWhitelistRepository) Delete(employeeId string) error {
	delete(inMemoryWhiteList, employeeId)
	return nil
}
