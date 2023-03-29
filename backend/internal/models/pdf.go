package models

import (
	"io"
	ioutil "io/ioutil"
	"mime/multipart"
	"os"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	// "github.com/lonrun/SmartAssist/backend/internal/consts"
)

type PDF struct {
	Content []byte
}

func NewPDF(content []byte) *PDF {
	return &PDF{Content: content}
}

func NewPDFFromFile(file *multipart.FileHeader) (*PDF, error) {
	// Open file
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read file content
	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	// Create PDF
	pdf := &PDF{Content: content}
	return pdf, nil
}

func (p *PDF) Summarize() (string, error) {
	// Create Maroto PDF instance
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	// Add text to PDF
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text("Summary of PDF Content", props.Text{
				Top:    20,
				Style:  consts.Bold,
				Family: consts.Courier,
				Size:   12,
			})
		})
	})
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text(string(p.Content), props.Text{})
		})
	})

	// Render PDF
	err := m.OutputFileAndClose(os.TempDir() + "/pdf_summary.pdf")
	if err != nil {
		return "", err
	}

	// Read PDF summary
	summary, err := ioutil.ReadFile(os.TempDir() + "/pdf_summary.pdf")
	if err != nil {
		return "", err
	}

	// Delete temporary PDF file
	err = os.Remove(os.TempDir() + "/pdf_summary.pdf")
	if err != nil {
		return "", err
	}

	// Return PDF summary
	return string(summary), nil
}
