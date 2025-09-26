package dreamer

import (
	"github.com/google/uuid"
)

type Entity struct {
	ID         EntityID
	Properties []Property
}

func NewEntity(properties []Property) *Entity {
	return &Entity{
		ID:         EntityID(uuid.NewString()),
		Properties: properties,
	}
}
