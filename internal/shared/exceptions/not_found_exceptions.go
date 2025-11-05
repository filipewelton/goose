package exceptions

import "errors"

type notFoundExceptions struct {
	ErrUserNotFound error
}

var NotFoundExceptions = notFoundExceptions{
	ErrUserNotFound: errors.New("usuário não encontrado"),
}
