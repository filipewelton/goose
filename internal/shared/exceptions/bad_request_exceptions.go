package exceptions

import "errors"

type badRequestExceptionNames struct {
	ErrInvalidCardNumber                error
	ErrInvalidName                      error
	ErrPasswordRequiresLowercaseLetter  error
	ErrPasswordRequiresUppercaseLetter  error
	ErrPasswordRequiresDigit            error
	ErrPasswordRequiresSpecialCharacter error
	ErrPasswordInvalidLength            error
}

var BadRequestExceptions = badRequestExceptionNames{
	ErrInvalidCardNumber: errors.New("número de matrícula inválido"),

	ErrInvalidName: errors.New("nome inválido"),

	ErrPasswordRequiresLowercaseLetter: errors.New("senha deve conter ao menos uma letra minúscula"),

	ErrPasswordRequiresUppercaseLetter: errors.New("senha deve conter ao menos uma letra maiúscula"),

	ErrPasswordRequiresDigit: errors.New("senha deve conter ao menos um dígito"),

	ErrPasswordRequiresSpecialCharacter: errors.New("senha deve conter ao menos um caractere especial"),

	ErrPasswordInvalidLength: errors.New("senha deve ter entre 8 e 48 caracteres"),
}
