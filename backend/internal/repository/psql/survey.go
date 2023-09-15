package psql

import (
	"mado/internal/core/survey"
	"mado/pkg/database/postgres"
)

// Survey is a Survey repository.
type SurveyrRepository struct {
	db *postgres.Postgres
}

// NewSurveyRepository creates a new UserRepository.
func NewSurveyrRepository(db *postgres.Postgres) SurveyrRepository {
	return SurveyrRepository{
		db: db,
	}
}

// todo implement this
func (s SurveyrRepository) Create(*survey.SurveyRequirements) (*survey.SurveyRequirements, error) {
	return nil, nil
}

