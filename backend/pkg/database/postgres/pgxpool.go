package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsolationLevel represents the transaction isolation level.
type IsolationLevel string

const (
	// ReadCommitted is the default isolation level and offers a balance of concurrency and consistency.
	ReadCommitted IsolationLevel = "READ COMMITTED"

	// Serializable offers the highest level of isolation but can lead to more contention.
	Serializable IsolationLevel = "SERIALIZABLE"
)

// PgxPool is a PostgreSQL connection pool.
// abstracts the PostgreSQL connection pool. It contains methods for executing queries and managing the pool.
type PgxPool interface { //we heed ti for implementating postgre interface
	Close()
	// Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row

	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)

	// BeginTx starts a new transaction with the given isolation level.
	// If the isolation level is empty, it uses the default (ReadCommitted).
	// BeginTx(ctx context.Context, isolationLevel IsolationLevel) (pgx.Tx, error)

	// Transaction runs a function within a transaction and automatically commits or rolls back based on the error.
	// The isolation level is optional, and it uses the default (ReadCommitted) if not specified.
	// Transaction(ctx context.Context, fn func(tx pgx.Tx) error, isolationLevel IsolationLevel) error
}
