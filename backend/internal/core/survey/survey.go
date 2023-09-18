package survey

import (
	"context"
	"time"
)

// Repository is a user repository.
type Repository interface {
	// Create(*survey.SurveyRequirements) (*survey.SurveyRequirements, error)
	Create(*SurveyRequirements, context.Context) (*SurveyRequirements, error)
}

// Service is a user service interface.
type Service struct {
	surveyRepository Repository
}

// NewService creates a new user service.
func NewService(surveyRepository Repository) Service {
	return Service{
		surveyRepository: surveyRepository,
	}
}

func (s Service) Create(requirements *SurveyRequirements) (*SurveyRequirements, error) {

	if err := s.ValidateSurveyRequirements(requirements); err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return s.surveyRepository.Create(requirements, ctx)
}

func (s Service) ValidateSurveyRequirements(requirements *SurveyRequirements) error {
	if requirements == nil {
		return ErrSurvey
	}

	if requirements.Name == "" {
		return ErrSurveyName
	}

	if len(requirements.Questions) == 0 {
		return ErrSurveyQuestion
	}

	// Add more checks for other fields as needed

	return nil
}
