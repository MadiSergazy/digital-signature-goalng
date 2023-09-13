package postgres

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxPoolSize  int           = 1
	defaultConnAttempts int           = 10
	defaultConnTimeout  time.Duration = 5 * time.Second
)

// ErrUnableToConnect is returned when unable to connect to the postgres.
var ErrUnableToConnect = errors.New("all attempts are exceeded. Unable to connect to instance")

// Postgres is a postgres connection.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	Builder      sq.StatementBuilderType
	Pool         PgxPool
}

// New creates a new postgres connection.
func New(ctx context.Context, connectionConfig ConnectionConfig, opts ...Option) (*Postgres, error) {
	instance := &Postgres{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	// Apply any custom options passed as arguments to the function. IT will work only if ...Option arg was passed example WithMaxPoolSize, WithConnAttempts and WithConnTimeout functions
	for _, opt := range opts {
		opt(instance)
	}

	// Set up the SQL query builder.
	instance.Builder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	dsn := connectionConfig.getDSN()
	fmt.Println("DSN: ", dsn)
	ctx, cancel := context.WithTimeout(context.Background(), defaultConnTimeout)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	// Ping the database to check the connection
	if err := dbpool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}

	instance.Pool = dbpool
	// defer dbpool.Close()

	/*var err error
	if err := DoWithTrials(func() error {
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		instance.Pool, err = pgxpool.ParseConfig() // Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, defaultConnAttempts, 5*time.Second); err != nil {
		log.Fatal("error do with postgresql")
	}*/

	/*
		poolCfg, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return nil, fmt.Errorf("unable to parse config: %w", err)
		}
		fmt.Println("poolCfg: ", poolCfg)

		poolCfg.MaxConns = int32(instance.maxPoolSize)

		for instance.connAttempts > 0 {
			instance.Pool, err = pgxpool.ConnectConfig(ctx, poolCfg)
			if err == nil {
				break
			}

			log.Printf("Postgres is trying to connect, attempts left: %d", instance.connAttempts)
			time.Sleep(instance.connTimeout)

			instance.connAttempts--
		}*/

	// if err != nil {
	// 	return nil, ErrUnableToConnect
	// }
	fmt.Println("POOL postgres: ", instance.Pool)
	return instance, nil
}
