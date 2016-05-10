package viewmodels

import (
	"GoWebApplication/src/models"
)

type Clients struct {
	Title   string
	Active  string
	Clients []models.Client
}

func GetClients() Clients {
	//TODO: get list of clients from the database

	result := Clients{
		Title:  "Clients",
		Active: "clients",
	}

	return result
}
