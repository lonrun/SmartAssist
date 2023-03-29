package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonrun/SmartAssist/backend/internal/app"
	"github.com/lonrun/SmartAssist/backend/internal/models"
)

type ticketRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Type        string `json:"type"`
}

func TicketHandler(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req ticketRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		ticket := &models.Ticket{
			Source:      req.Source,
			Destination: req.Destination,
			Date:        req.Date,
			Type:        req.Type,
		}

		response, err := a.BookTicket(ticket)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to book ticket"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": response})
	}
}
