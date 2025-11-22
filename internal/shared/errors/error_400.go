package errors

import "fmt"

var Err400 = struct {
	ErrInvalidUserName                     error
	ErrInvalidEmployeeId                   error
	ErrLowercaseLettersIsMissingInPassword error
	ErrUppercaseLettersIsMissingInPassword error
	ErrNumbersIsMissingInPassword          error
	ErrPasswordTooShort                    error
	ErrPasswordTooLong                     error
	ErrFailedToReadTheRequestBody          error
}{
	ErrInvalidUserName: fmt.Errorf("nome de usuário é inválido"),

	ErrInvalidEmployeeId: fmt.Errorf("número de crachá inválido"),

	ErrLowercaseLettersIsMissingInPassword: fmt.Errorf("a senha deve conter letras minusculas"),

	ErrUppercaseLettersIsMissingInPassword: fmt.Errorf("a senha deve conter letras maiúsculas"),

	ErrNumbersIsMissingInPassword: fmt.Errorf("a senha deve conter números"),

	ErrPasswordTooShort: fmt.Errorf("a senha deve conter no mínimo 8 caracteres"),

	ErrPasswordTooLong: fmt.Errorf("a senha deve conter nom máximo 48 caracteres"),

	ErrFailedToReadTheRequestBody: fmt.Errorf("falha ao ler o corpo da requisição"),
}
