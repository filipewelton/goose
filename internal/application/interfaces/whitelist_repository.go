package interfaces

type WhitelistRepository interface {
	Insert(string) error
	Has(string) (bool, error)
	Delete(string) error
}
