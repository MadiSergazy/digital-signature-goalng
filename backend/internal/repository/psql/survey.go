package psql

import (
	"log"
	"mado/internal/core/survey"
	"mado/pkg/database/postgres"

	"github.com/gin-gonic/gin"
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

func (s SurveyrRepository) GetSurviesByUserID(user_iin string, ctx *gin.Context) (response *survey.SurveyResponse, err error) {
	query := "SELECT * FROM survey WHERE iin = $1"
	rows, err := s.db.Pool.Query(ctx, query, user_iin)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()
	var surveyResponse *survey.SurveyResponse
	for rows.Next() {
		// Scan the row into variables
		if err := rows.Scan(&surveyResponse.ID, &surveyResponse.Name, &surveyResponse.Status, &surveyResponse.Rka, &surveyResponse.Rc_name, &surveyResponse.Adress, &surveyResponse.Question_id, &surveyResponse.CreatedAt, &surveyResponse.User_id); err != nil {
			log.Fatalf("Error scanning row: %v", err)
		}
	}
	return surveyResponse, nil
}
