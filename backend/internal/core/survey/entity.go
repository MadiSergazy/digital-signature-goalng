package survey

import "errors"

var (
	ErrSurvey         = errors.New("surveyRequirements is nil")
	ErrSurveyName     = errors.New("name field is empty")
	ErrSurveyQuestion = errors.New("questions field is empty")
	ErrAlreadyExist   = errors.New("survey with given email or nickname already exist")
	ErrNotFound       = errors.New("survey not found")
)

type SurveyRequirements struct {
	ID         string     `json:"id,omitempty"`
	UserID     int        `json:"user_id,omitempty"`
	Name       string     `json:"name"`
	Rka        string     `json:"rka,omitempty"`
	RcName     string     `json:"rc_name,omitempty"`
	Adress     string     `json:"address,omitempty"`
	Questions  []Question `json:"questions"`
	CreateDate string     `json:"create_date,omitempty"`
}

type Question struct {
	// Name string `json:"name"`
	Description string `json:"description"`
} //return []question id
