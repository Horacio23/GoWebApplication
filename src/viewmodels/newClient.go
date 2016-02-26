package viewmodels

import (

)

type NewClient struct {
	Title string
	Active string
	Client Client
}

func GetNewClient() NewClient {
	result := NewClient {
		Title: "New Client",
		Active: "newClient",
//		Client: Client{
//			Name: "Horacio Delgado",
//			Address: "1353 nw 123 street Miami Fl",
//			Phone: "788-895-8544",
//		},
	}
	
	return result
}