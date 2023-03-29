package app

import (
	"github.com/lonrun/SmartAssist/backend/internal/models"
	"github.com/lonrun/SmartAssist/backend/internal/services"
)

type App struct {
	pdfService    *services.PDFService
	ticketService *services.TicketService
	deviceService *services.DeviceService
}

func NewApp(pdfService *services.PDFService, ticketService *services.TicketService, deviceService *services.DeviceService) *App {
	return &App{
		pdfService:    pdfService,
		ticketService: ticketService,
		deviceService: deviceService,
	}
}

func (a *App) SummarizePDF(pdf *models.PDF) (string, error) {
	return a.pdfService.Summarize(pdf.Content)
}

func (a *App) BookTicket(ticket *models.Ticket) (string, error) {
	return a.ticketService.Book(ticket)
}

func (a *App) ControlDevice(device *models.Device) (string, error) {
	return a.deviceService.Control(device)
}
