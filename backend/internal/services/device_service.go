package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/lonrun/SmartAssist/backend/internal/models"
	"github.com/lonrun/SmartAssist/backend/pkg/database"
)

type DeviceService struct{}

func NewDeviceService(db *database.Database) *DeviceService {
	return &DeviceService{}
}

func (s *DeviceService) Control(device *models.Device) (string, error) {
	// Simulate device control
	time.Sleep(1 * time.Second)
	switch device.Action {
	case "on":
		return fmt.Sprintf("Device %s turned on", device.Name), nil
	case "off":
		return fmt.Sprintf("Device %s turned off", device.Name), nil
	default:
		return "", errors.New("Invalid device action")
	}
}
