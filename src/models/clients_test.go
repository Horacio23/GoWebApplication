package models

import (
	"fmt"
	"testing"
)

func Test_InsertingNewClient(t *testing.T) {

	client := Client{}
	client.FirstName = "firstName"
	client.LastName = "lastName"
	client.Address = "address"
	client.City = "city"
	client.State = "FL"
	client.Zip = "78954"
	client.Phone = "7896545568"
	client.EntranceDate = "entranceDate"
	client.LastTransaction = "05/17/2016"
	client.TransactionDate = "05/17/2016"
	client.Notes = "notes"

	client, err := CreateClient(client)
	if err == nil {
		fmt.Println(client)
	} else {
		t.Log("Test for creating new client failed: " + err.Error())

		t.FailNow()
	}
}

func Test_GetClient(t *testing.T) {
	client, err := GetClient(3)
	if err == nil {
		fmt.Println(client)
		t.Log(client.FirstName)
	} else {
		t.Log("Getting a client failed")
		t.FailNow()
	}
}

func Test_CheckDate(t *testing.T) {
	clients, err := CheckDates()

	if err == nil {
		fmt.Println("the clients are", clients)
		t.Log("Client array lengths")
		t.Log(len(clients))
	} else {
		t.Log("Getting a client failed:", err.Error())
		t.FailNow()
	}
}
