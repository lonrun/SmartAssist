package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonrun/SmartAssist/backend/internal/app"
	"github.com/lonrun/SmartAssist/backend/internal/models"
)

func PDFHandler(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get file from request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// Open file and create PDF model
		pdf, err := models.NewPDFFromFile(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse PDF"})
			return
		}

		// Summarize PDF
		summary, err := a.SummarizePDF(pdf)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to summarize PDF"})
			return
		}

		// Return summary
		c.JSON(http.StatusOK, gin.H{"summary": summary})
	}
}
