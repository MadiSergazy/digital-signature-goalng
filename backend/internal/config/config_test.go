package config_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"mado/internal/config"
)

type env struct {
	environment      string
	httpHost         string
	httpPort         string
	postgresHost     string
	postgresPort     string
	postgresDBName   string
	postgresUser     string
	postgresPassword string
	baseURL          string
	// tokenSecretKey   string
	corsAllowOrigins string
}

func setEnv(t *testing.T, env env) {
	t.Helper()

	require.NoError(t, os.Setenv("ENVIRONMENT", env.environment))
	require.NoError(t, os.Setenv("HTTP_HOST", env.httpHost))
	require.NoError(t, os.Setenv("HTTP_PORT", env.httpPort))
	require.NoError(t, os.Setenv("POSTGRES_HOST", env.postgresHost))
	require.NoError(t, os.Setenv("POSTGRES_PORT", env.postgresPort))
	require.NoError(t, os.Setenv("POSTGRES_DBNAME", env.postgresDBName))
	require.NoError(t, os.Setenv("POSTGRES_USER", env.postgresUser))
	require.NoError(t, os.Setenv("POSTGRES_PASSWORD", env.postgresPassword))
	require.NoError(t, os.Setenv("BASE_URL", env.baseURL))
	// require.NoError(t, os.Setenv("TOKEN_SECRET_KEY", env.tokenSecretKey))
	require.NoError(t, os.Setenv("CORS_ALLOW_ORIGINS", env.corsAllowOrigins))
}

func TestGet(t *testing.T) {
	t.Parallel()

	env := env{
		environment:      "test",
		httpHost:         "192.168.1.32",
		httpPort:         "8080",
		postgresHost:     "postgres",
		postgresPort:     "5431",
		postgresDBName:   "test_conduit",
		postgresUser:     "test_conduit",
		postgresPassword: "test",
		baseURL:          "https://sigex.kz",
		// tokenSecretKey:   "secret",
		corsAllowOrigins: "*",
	}

	want := &config.Config{
		Environment: "test",
		HTTP: config.HTTP{
			Host:           "192.168.1.32",
			Port:           "8080",
			MaxHeaderBytes: 1,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
		Postgres: config.Postgres{
			Host:     "postgres",
			Port:     "5431",
			DBName:   "test_conduit",
			User:     "test_conduit",
			Password: "test",
			SSLMode:  "disable",
		},
		Logger: config.Logger{
			Level: "info",
		},
		SigexEndpoints: config.SigexEndpoints{
			BaseUrl: "https://sigex.kz",
		},
		// Token: config.Token{
		// 	SecretKey: "secret",
		// 	Expired:   15 * time.Minute,
		// },
		CORS: config.CORS{
			AllowOrigins: []string{"http://localhost:3000"},
		},
	}

	setEnv(t, env)

	got := config.Get()
	require.True(t, reflect.DeepEqual(want, got))
}
