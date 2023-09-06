package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username *string
	Email    *string
	IIN      *string
	BIN      *string
}
