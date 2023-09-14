package survey

type SurveyRequirements struct {
	Name      string     `json:"name"`
	Rka       string     `json:"rka"`
	RcName    string     `json:"rc_name"`
	Adress    string     `json:"address"`
	Questions []Question `json:"questions"`
}

type Question struct {
	Name string `json:"name"`
	Text string `json:"text"`
}
