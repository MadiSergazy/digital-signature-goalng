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

	// Ensure you have a valid database connection
	if ur.db == nil {
		fmt.Println("database connection is nil")
		return nil, errors.New("database connection is nil")
	}

	// Prepare the SQL statement
	sqlStatement := `
		INSERT INTO users (iin, email, bin, name, is_manager) 
		VALUES ($1, $2, $3, $4, $5);
	`

	logger.FromContext(ctx).Debug("create user query", zap.String("sql", sqlStatement), zap.Any("args", dto))

	// Execute the SQL statement
	result, err := ur.db.Pool.Exec(ctx, sqlStatement, dto.IIN, dto.Email, dto.BIN, dto.Username, false)
	if err != nil {
		fmt.Println("error executing sql statement")
		return nil, fmt.Errorf("%w%w", errInsert, err)
	}

	// Check the number of rows affected (usually for error checking)
	rowsAffected := result.RowsAffected()
	if rowsAffected != 1 {
		fmt.Println("number of rows affected err")
		return nil, fmt.Errorf("expected 1 row to be affected, but %d rows were affected", rowsAffected)
	}

	// Optionally, you can retrieve the newly inserted user if your database supports returning the inserted row.
	// Otherwise, you may want to fetch the user by some unique identifier (e.g., ID) and return it here.

	return dto, nil
}

func (ur UserRepository) GetAllRows(ctx context.Context) ([]*user.User, error) {
	if ur.db == nil {
		return nil, errors.New("database connection is nil")
	}

	// Prepare the SQL statement
	sqlStatement := `SELECT * FROM users`

	// Execute the SQL statement and retrieve the result set
	rows, err := ur.db.Pool.Query(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*user.User

	// Iterate through the result set and scan each row into a user struct
	for rows.Next() {
		var u user.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.BIN, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
