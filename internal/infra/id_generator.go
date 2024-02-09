package infra

import "github.com/google/uuid"

type UuidGenerator struct{}

func (u *UuidGenerator) Generate() uuid.UUID {
	return uuid.New()
}
