package viewmodels

import (
	"models"
)

type UpdateClientView struct {
	Title string
	Active string
	Client models.Client
}

func  GetUpdateClient(id int) UpdateClientView{
	client := models.GetClient(id)
	result:= UpdateClientView{
		Title: "Client Update Page",
		Active: "update",
		Client: client,
	}
	
	return result
}



