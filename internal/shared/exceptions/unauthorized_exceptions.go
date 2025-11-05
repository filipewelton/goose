package exceptions

import "errors"

type unauthorizedExceptions struct {
	ErrAPITokenExpired error
	ErrInvalidAPIToken error
}

var UnauthorizedExceptions = unauthorizedExceptions{
	ErrAPITokenExpired: errors.New("token de API está expirado"),
	ErrInvalidAPIToken: errors.New("token de API é inválido"),
}
