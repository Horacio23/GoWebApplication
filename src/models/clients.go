package models

import (

)

type Client struct {
	Name string
	Address string
	Phone string
	Id int
}

func (this *Client) SetName(name string) {
	this.Name = name
}

func (this *Client) SetAddress(address string) {
	this.Address = address
}

func (this *Client) SetPhone(phone string) {
	this.Phone = phone
}

func (this *Client) SetId(id int) {
	this.Id = id
}

func GetClients() []Client {
	//TODO: get list of clients from the database
	result := []Client{
		Client{
			Name: "horacio",
			Address: "555 nw 19th, Miami, Fl, 32182",
			Phone: "786-555-4785",
			Id:1,
		},
	}
	
	return result
}

func GetClient(id int) Client {
	result := Client{}

	if id == 1 {
		result = Client{
			Name: "Geraldo",
			Address: "555 nw 19th, Miami, Fl, 32182",
			Phone: "786-555-4785",
			Id:1,
		}
	}

	
	return result
	
}
