package viewmodels

import (
	"models"
)

type ClientView struct {
	Title string
	Active string
	Client models.Client
}

func GetClient(id int) ClientView {
	result := ClientView{
		Title: "Client",
		Active: "clients",
	}

	result.Client = models.GetClient(id)

	return result
	
}