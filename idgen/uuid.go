package idgen

import (
	"github.com/google/uuid"
)

type UUIDGen struct{}

func (u *UUIDGen) NextID() string {
	return uuid.NewString()
}
