package petition

import (
	"context"
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

type Repository interface {
	Save(ctx context.Context, dto *PetitionData) (*PetitionData, error)
}

// Service is a user service interface.
type Service struct {
	petitionRepository Repository
	logger             *zap.Logger
}

// NewService creates a new user service.
func NewService(petitionRepository Repository, logger *zap.Logger) Service {
	return Service{
		petitionRepository: petitionRepository,
		logger:             logger,
	}
}

// Todo save to db ?
func (s *Service) GeneratePetitionPDF(data *PetitionData) error {

	// Create a new template
	tmpl, err := template.New("petition").Parse(TemplateHTML)
	if err != nil {
		return (err)
	}

	// Create a buffer to hold the generated HTML content
	var htmlContentBuffer strings.Builder
	if err := tmpl.Execute(&htmlContentBuffer, data); err != nil {
		return (err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContentBuffer.String()))
	if err != nil {
		fmt.Println("Here is an error: ", err)
		return (err)
	}

	// Save the modified HTML to a temporary file
	tempHTMLFileName := "temp.html"
	tempHTMLFile, err := os.Create(tempHTMLFileName)
	if err != nil {
		return (err)
	}
	defer tempHTMLFile.Close()

	// Use goquery to write the modified HTML to the file
	htmlContent, err := doc.Html()
	if err != nil {
		return (err)
	}
	_, err = tempHTMLFile.WriteString(htmlContent)
	if err != nil {
		return (err)
	}

	// Generate PDF using wkhtmltopdf
	//todo change namingOfFile
	pdfFileName := "output.pdf"
	cmd := exec.Command("wkhtmltopdf", tempHTMLFileName, pdfFileName)
	err = cmd.Run()
	if err != nil {
		return (err)
	}

	fmt.Printf("PDF generated: %s\n", pdfFileName)
	return nil
}
