package psql

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"mado/internal/core/user"
	"mado/pkg/database/postgres"
	"mado/pkg/logger"
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

var errInsert = errors.New("can not insert user: ")

// TODO do it properly
func (ur UserRepository) Create(ctx context.Context, dto *user.User) (*user.User, error) {

	sql, args, err := ur.db.Builder.Insert("users").Columns("IIN", "email", "BIN").Values(dto.IIN, dto.Email, dto.BIN).ToSql()
	if err != nil {
		return &user.User{}, fmt.Errorf("can not build insert user query: %w", err)
	}

	logger.FromContext(ctx).Debug("create user query", zap.String("sql", sql), zap.Any("args", args))

	tag, err := ur.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return &user.User{}, fmt.Errorf("%w%w", errInsert, err)
	}
	if !tag.Insert() {
		return &user.User{}, fmt.Errorf("%w%w", errInsert, err)
	}
	if rows := tag.RowsAffected(); rows != 1 {
		return &user.User{}, fmt.Errorf("%w%w", errInsert, err)
	}

	return nil, nil
}
