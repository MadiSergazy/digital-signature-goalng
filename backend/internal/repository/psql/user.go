package psql

import (
	"context"

	"mado/internal/core/user"
	"mado/pkg/database/postgres"
)

// UserRepository is a user repository.
type UserRepository struct {
	db *postgres.Postgres
}

// NewUserRepository creates a new UserRepository.
func NewUserRepository(db *postgres.Postgres) UserRepository {
	return UserRepository{
		db: db,
	}
}

// TODO do it properly
func (ur UserRepository) Create(ctx context.Context, dto *user.User) (*user.User, error) {
	return nil, nil
}
