package viewmodels

import (

)

type Clients struct {
	Title string
	Active string
	Clients []Client
}

type Client struct {
	Name string
	Address string
	Phone string
	Id int
}

func GetClients() Clients {
	//TODO: get list of clients from the database
	
	result := Clients{
		Title: "Clients",
		Active: "clients",
	}
	
	return result
}

func GetClient(id int) Clients {
	result := Clients{
		Title: "Client",
		Active: "clients",
	}

	if id == 1 {
		client1 := Client{
			Name: "Geraldo",
			Address: "555 nw 19th, Miami, Fl, 32182",
			Phone: "786-555-4785",
			Id:1,
		}
		client2 := Client{
			Name: "Geraldo",
			Address: "555 nw 19th, Miami, Fl, 32182",
			Phone: "786-555-4785",
			Id:2,
		}
		client3 := Client{
			Name: "Geraldo",
			Address: "555 nw 19th, Miami, Fl, 32182",
			Phone: "786-555-4785",
			Id:3,
		}
		
		result.Clients = []Client{
			client1,
			client2,
			client3,
		}
	}

	return result
	
}


















