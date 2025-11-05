package exceptions

import "errors"

type conflictExceptions struct {
	ErrCardNumberAlreadyRegistered error
}

var ConflictExceptions = conflictExceptions{
	ErrCardNumberAlreadyRegistered: errors.New("número de matrícula já cadastrado"),
}
