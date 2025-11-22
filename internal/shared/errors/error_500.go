package errors

import "fmt"

var Error500 = struct {
	ErrMissingEnvironmentVariable      error
	ErrPostgresConnectionFailure       error
	ErrPostgresDisconnectionFailure    error
	ErrRedisConnectionFailure          error
	ErrPasswordHashGenerationFailure   error
	ErrUserDeletionFailedInPostgreSQL  error
	ErrKeyInsertionFailedInRedis       error
	ErrAPITokenGenerationFailure       error
	ErrUserCredentialGenerationFailure error
	ErrUndefinedUserRepository         error
}{
	ErrMissingEnvironmentVariable: fmt.Errorf("alguma variável de ambiente está faltando"),

	ErrPostgresConnectionFailure: fmt.Errorf("falha de conexão com o PostgreSQL"),

	ErrPostgresDisconnectionFailure: fmt.Errorf("falha ao desconectar o PostgreSQL"),

	ErrRedisConnectionFailure: fmt.Errorf("falha ao conectar com Redis"),

	ErrPasswordHashGenerationFailure: fmt.Errorf("falha na geração do hash da senha"),

	ErrUserDeletionFailedInPostgreSQL: fmt.Errorf("falha na exclusão de usuário no PostgreSQL"),

	ErrKeyInsertionFailedInRedis: fmt.Errorf("falha na inserção de chave no Redis"),

	ErrAPITokenGenerationFailure: fmt.Errorf("falha na geração do token de API"),

	ErrUserCredentialGenerationFailure: fmt.Errorf("falha na geração das credenciais de usuário"),

	ErrUndefinedUserRepository: fmt.Errorf("repositório de usuário indefinido"),
}
