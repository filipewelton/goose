package valueobjects

import "github.com/oklog/ulid/v2"

type EntityId string

func (e *EntityId) Generate() {
	var id string = ulid.Make().String()
	*e = EntityId(id)
}

func (e *EntityId) Set(id string) {
	*e = EntityId(id)
}

func (e *EntityId) Get() string {
	return string(*e)
}
