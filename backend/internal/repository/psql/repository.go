package psql

import "mado/pkg/database/postgres"

// Repositories is a collection of all repositories in the system.
type Repositories struct {
	User   UserRepository
	Survey SurveyrRepository
}

// NewRepositories returns a new instance of Repositories.
func NewRepositories(db *postgres.Postgres) Repositories {
	return Repositories{
		User:   NewUserRepository(db),
		Survey: NewSurveyrRepository(db),
	}
}
