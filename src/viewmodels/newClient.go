package viewmodels

import (
	"models"

)

type NewClient struct {
	Title string
	Active string
	Client models.Client
}

func GetNewClient() NewClient {
	
	result := NewClient {
		Title: "New Client",
		Active: "newClient",
	}
	
	return result
}