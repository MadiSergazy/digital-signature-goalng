package api

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	"mado/internal/config"
	"mado/pkg/database/postgres"
)

// App is a application interface.
type App struct {
	logger     *zap.Logger
	db         *postgres.Postgres
	httpServer httpserver.Server
}

// New creates a new App.
func New(ctx context.Context, logger *zap.Logger) (App, error) {
	cfg := config.Get()

	postgresInstance, err := postgres.New(
		ctx,
		postgres.NewConnectionConfig(
			cfg.Postgres.Host,
			cfg.Postgres.Port,
			cfg.Postgres.DBName,
			cfg.Postgres.User,
			cfg.Postgres.Password,
			cfg.Postgres.SSLMode,
		),
	)
	if err != nil {
		return App{}, fmt.Errorf("can not connect to postgres: %w", err)
	}

	// passwordHasher := hash.NewArgon2Hasher()

	// tokenMaker, err := token.NewJWTMaker(cfg.Token.SecretKey)
	// if err != nil {
	// 	return App{}, fmt.Errorf("failed to create token maker: %w", err)
	// }

	repositories := psql.NewRepositories(postgresInstance)
	services := domain.NewServices(repositories, passwordHasher)

	router := httphandler.NewRouter(httphandler.Deps{
		TokenMaker: tokenMaker,
		Logger:     logger,
		Services:   services,
	})

	return App{
		logger: logger,
		db:     postgresInstance,
		httpServer: httpserver.New(
			router,
			httpserver.WithHost(cfg.HTTP.Host),
			httpserver.WithPort(cfg.HTTP.Port),
			httpserver.WithMaxHeaderBytes(cfg.HTTP.MaxHeaderBytes),
			httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
			httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		),
	}, nil
}
