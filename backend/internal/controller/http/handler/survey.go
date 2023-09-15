package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"mado/internal/core/survey"
)

type SurveyService interface {
	// Create(ctx context.Context, dto user.CreateDTO) (user.User, error)
	Create(*survey.SurveyRequirements) (*survey.SurveyRequirements, error)
	// GetAllRows()()
}

type surveyDeps struct {
	router *gin.RouterGroup

	surveyService SurveyService
}

type surveyHandler struct {
	surveyService SurveyService
}

func newSurveyHandler(deps surveyDeps) {
	handler := surveyHandler{
		surveyService: deps.surveyService,
	}

	usersGroup := deps.router.Group("/survey")
	{

		usersGroup.POST("/create", handler.CreateSurvey)
	}

}

func (h surveyHandler) CreateSurvey(c *gin.Context) {
	var request *survey.SurveyRequirements
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

}
