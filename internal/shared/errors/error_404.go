package errors

import "fmt"

var Err404 = struct {
	ErrUserNotFound       error
	ErrEmployeeIdNotFound error
}{
	ErrUserNotFound: fmt.Errorf("usuário não encontrado"),

	ErrEmployeeIdNotFound: fmt.Errorf("número de crachá não encontrado na lista de permissões"),
}
