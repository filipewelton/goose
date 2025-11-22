package errors

import "fmt"

var Error401 = struct {
	ErrInvalidUserCredential    error
	ErrUnauthenticatedAPIAccess error
	ErrInvalidUserCredentials   error
	ErrUserUnauthenticated      error
}{
	ErrInvalidUserCredential: fmt.Errorf("crachá ou senha inválidos"),

	ErrUnauthenticatedAPIAccess: fmt.Errorf("acesso à API não autorizado"),

	ErrInvalidUserCredentials: fmt.Errorf("credenciais de usuário inválidas"),

	ErrUserUnauthenticated: fmt.Errorf("usuário não autenticado"),
}
