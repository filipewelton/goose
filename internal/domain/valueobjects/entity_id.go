package valueobjects

import (
	"retail_flow/internal/shared/exceptions"
	"retail_flow/internal/shared/lib"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type EntityID struct {
	value string
}

func (e *EntityID) New() error {
	var log = lib.Logger()

	defer log.Sync()

	uuid, err := uuid.NewV7()

	if err != nil {
		log.Error(
			exceptions.InternalExceptions.ErrEntityIDGenerationFailed.Error(),
			zap.Error(err),
		)

		return exceptions.InternalExceptions.ErrEntityIDGenerationFailed
	}

	e.value = uuid.String()
	return nil
}

func (e *EntityID) GetValue() string {
	return e.value
}

func (e *EntityID) SetValue(value string) {
	e.value = value
}
