package services

import (
	"fmt"

	"github.com/lonrun/SmartAssist/backend/internal/models"
	"github.com/lonrun/SmartAssist/backend/pkg/database"
)

type TicketService struct{}

func NewTicketService(db *database.Database) *TicketService {
	return &TicketService{}
}

func (s *TicketService) Book(ticket *models.Ticket) (string, error) {
	// Book ticket
	confirmation := fmt.Sprintf("Ticket booked for %s to %s on %s (%s)", ticket.Source, ticket.Destination, ticket.Date, ticket.Type)
	return confirmation, nil
}
