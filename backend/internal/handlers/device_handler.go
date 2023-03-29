package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonrun/SmartAssist/backend/internal/app"
	"github.com/lonrun/SmartAssist/backend/internal/models"
)

type deviceRequest struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

func DeviceHandler(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req deviceRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		device := &models.Device{Name: req.Name, Action: req.Action}
		response, err := a.ControlDevice(device)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to control device"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": response})
	}
}
