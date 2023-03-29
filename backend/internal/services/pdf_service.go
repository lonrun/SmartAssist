package services

import (
	"github.com/lonrun/SmartAssist/backend/internal/models"
	"github.com/lonrun/SmartAssist/backend/pkg/database"
)

type PDFService struct{}

func NewPDFService(db *database.Database) *PDFService {
	return &PDFService{}
}

func (s *PDFService) Summarize(content []byte) (string, error) {
	pdf := models.NewPDF(content)
	return pdf.Summarize()
}
