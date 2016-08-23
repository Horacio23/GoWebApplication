package viewmodels

import (
	"models"
)

type Clients struct {
	Title   string
	Active  string
	User    string
	Clients []models.Client
}

func GetClients() Clients {
	result := Clients{
		Title:  "Clients",
		Active: "clients",
	}

	return result
}
