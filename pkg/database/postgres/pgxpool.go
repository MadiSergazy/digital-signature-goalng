package postgres

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// PgxPool is a PostgreSQL connection pool.
// abstracts the PostgreSQL connection pool. It contains methods for executing queries and managing the pool.
type PgxPool interface { //we heed ti for implementating postgre interface
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, optionsAndArgs ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
}
