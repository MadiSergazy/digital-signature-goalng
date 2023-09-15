package survey

// Repository is a user repository.
type Repository interface {
	// Create(*survey.SurveyRequirements) (*survey.SurveyRequirements, error)
	Create(*SurveyRequirements) (*SurveyRequirements, error)
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
	return nil, nil
}
