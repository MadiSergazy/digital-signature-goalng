package core

import (
	"mado/internal/core/user"
	"mado/internal/repository/psql"
)

// Services is a collection of all services in the system.
type Services struct {
	User user.Service
}

// NewServices returns a new instance of Services.
func NewServices(repositories psql.Repositories) Services {
	return Services{
		User: user.NewService(repositories.User),
	}
}
