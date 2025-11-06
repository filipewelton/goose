package exceptions

import "errors"

type internalExceptions struct {
	ErrFailureInParsingTheRequestBody     error
	ErrPostgresConnectionFailure          error
	ErrConnectionToPostgresNotEstablished error
	ErrFailedToInsertUserIntoDatabase     error
	ErrEnvironmentVariableIsMissing       error
	ErrEntityIDGenerationFailed           error
	ErrPasswordHashingFailed              error
	ErrJWTGenerationFailed                error
	ErrRedisConnectionFailure             error
	ErrRedisDisconnectionFailure          error
	ErrRedisInsertionFailed               error
	ErrRedisDeletionFailed                error
}

var InternalExceptions = internalExceptions{
	ErrFailureInParsingTheRequestBody: errors.New("falha na leitura do corpo da requisição"),

	ErrPostgresConnectionFailure: errors.New("falha ao conectar com Postgres"),

	ErrConnectionToPostgresNotEstablished: errors.New("conexão com Postgres não estabelecida"),

	ErrFailedToInsertUserIntoDatabase: errors.New("falha ao inserir usuário no banco de dados"),

	ErrEnvironmentVariableIsMissing: errors.New("alguma variável de ambiente está faltando"),

	ErrEntityIDGenerationFailed: errors.New("falha ao gerar o ID da entidade"),

	ErrPasswordHashingFailed: errors.New("falha ao gerar o hash da senha"),

	ErrJWTGenerationFailed: errors.New("falha ao gerar o JWT"),

	ErrRedisConnectionFailure: errors.New("falha ao conectar com Redis"),

	ErrRedisDisconnectionFailure: errors.New("falha ao desconectar o Redis"),

	ErrRedisInsertionFailed: errors.New("falha ao inserir uma chave no Redis"),

	ErrRedisDeletionFailed: errors.New("falha ao excluir uma chave no Redis"),
}
