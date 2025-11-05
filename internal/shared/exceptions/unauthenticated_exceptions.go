package exceptions

import "errors"

type unauthenticatedExceptions struct {
	ErrUnauthenticatedFrontend error
	ErrInvalidAPIToken         error
}

var UnauthenticatedExceptions = unauthenticatedExceptions{
	ErrUnauthenticatedFrontend: errors.New("frontend não autenticado"),
	ErrInvalidAPIToken:         errors.New("token de acesso à API é inválido"),
}
