package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"mado/internal/core/petition"
	"mado/pkg/errs"
)

type PetitionService interface {
	GetPetitionPdfByID(doc_id string) (response *petition.PetitionData, err error)
}

type petitionDeps struct {
	router          *gin.RouterGroup
	petitionService PetitionService
}

type petitionHandler struct {
	petitionService PetitionService
}

func newPetitionHandler(deps petitionDeps) {
	handler := petitionHandler{
		petitionService: deps.petitionService,
	}

	usersGroup := deps.router.Group("/petition")
	{
		usersGroup.GET("/download-pdf/:id", handler.GetPetitionPDF)
	}

}

func (h petitionHandler) GetPetitionPDF(c *gin.Context) {
	response, err := h.petitionService.GetPetitionPdfByID(c.Param("id"))
	if err != nil {
		if errors.Is(err, errs.ErrInvalidID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		} else if errors.Is(err, errs.ErrPdfFileNotFound) {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	// Set the appropriate headers for file download
	c.Header("Content-Disposition", strings.Join([]string{"attachment; filename", response.FileName + ".pdf"}, "="))
	// c.Header("Content-Disposition", "attachment; filename=your-file-name.pdf")
	c.Header("Content-Type", "application/pdf")
	c.Data(http.StatusOK, "application/pdf", response.PdfData)
}
