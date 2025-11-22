package user

type WhitelistRepository interface {
	Insert(string) error
	Has(string) (bool, error)
	Delete(string) error
}

type UserRepository interface {
	Insert(UserEntity) error
	FindByEmployeeId(string) (UserEntity, error)
	FindById(string) (UserEntity, error)
	DeleteByEmployeeId(string) error
}
