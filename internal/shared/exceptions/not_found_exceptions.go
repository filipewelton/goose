package exceptions

import "errors"

type notFoundExceptions struct {
	ErrUserNotFound        error
	ErrPreRegisterNotFound error
}

var NotFoundExceptions = notFoundExceptions{
	ErrUserNotFound: errors.New("usuário não encontrado"),

	ErrPreRegisterNotFound: errors.New("usuário não encontrado na lista de permissões"),
}
