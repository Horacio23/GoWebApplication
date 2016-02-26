package viewmodels

import (
	"models"
)

type Clients struct {
	Title string
	Active string
	Clients []models.Client
}


func GetClients() Clients {
	//TODO: get list of clients from the database
	
	result := Clients{
		Title: "Clients",
		Active: "clients",
	}
	result.Clients = models.GetClients()
	return result
}




















