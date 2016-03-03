package models

import (
	"testing"
	"fmt"
)



func Test_InsertingNewClient(t *testing.T) {
	client, err := CreateClient("Horacio","Delgado", "1235 sd 123 st", "Miami", "FL", "33192", "789-234-2341", "2013-02-02","2013-02-23")
	if err == nil {
		fmt.Println(client)
	}else{
		t.Log("Test for creating new client failed: "+err.Error())
		
		t.FailNow()
	}
}

func Test_GetClient(t *testing.T) {
	client, err := GetClient(3)
	if err == nil{
		fmt.Println(client)
		t.Log(client.FirstName)
	}else{
		t.Log("Getting a client failed")
		t.FailNow()
	}
}

