package models

type Ticket struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Type        string `json:"type"`
}
