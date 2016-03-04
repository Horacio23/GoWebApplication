package viewmodels

import (
	"models"
)

type UpdateClientView struct {
	Title string
	Active string
	Client models.Client
}

func  GetUpdateClient() UpdateClientView{

	result:= UpdateClientView{
		Title: "Client Update Page",
		Active: "update",
		
	}
	
	return result
}



