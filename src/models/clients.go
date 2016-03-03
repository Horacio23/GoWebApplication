package models

import (
	"errors"
	"fmt"
	"time"
)

type Client struct {
	Id int
	FirstName string
	LastName string
	Address string
	City string
	State string
	Zip string
	Phone string
	EntranceDate string
	TransactionDate string
}



func GetClients() []Client {
	//TODO: get list of clients from the database
	result := []Client{
	}
	
	return result
}

func GetClient(id int) (Client, error) {
	result := Client{}
	
	var entranceDate time.Time 
	var transactionDate time.Time
	
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		dbErr := db.QueryRow(`SELECT id, first_name, last_name, address, city, state, zip, phone, 
       entrance_date, transaction_date FROM clients where id=$1`,id).Scan(&result.Id, &result.FirstName, &result.LastName, &result.Address, &result.City, &result.State, &result.Zip, &result.Phone, &entranceDate, &transactionDate)
		
		result.EntranceDate = entranceDate.String()
		result.TransactionDate = transactionDate.String()
		
		fmt.Println("result id", id, "result",result)
		fmt.Println(dbErr)
		if dbErr == nil {
			fmt.Println("no error")
			return result, nil
		}else{
			fmt.Println("error in query")
			return Client{}, errors.New("Unable to create Client in the database: "+err.Error())
		}
	}else{
		fmt.Println("error with the db")
		return result, errors.New("Unable to get a database connection to save the session")
	}
	

	
	return result, nil
	
}

func CreateClient(firstName string, lastName string, address string, city string, state string, zip string, phone string, entranceDate string, transactionDate string) (Client, error) {
	result := Client{}
	result.FirstName = firstName
	result.LastName = lastName
	result.Address = address
	result.City = city
	result.Zip = zip
	result.State = state
	result.Phone = phone
	result.EntranceDate = entranceDate
	result.TransactionDate = transactionDate
	
	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO clients
			(first_name, last_name, address, city, state, zip, phone, entrance_date, transaction_date)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			RETURNING id`, firstName, lastName, address, city, state, zip, phone, entranceDate, transactionDate).Scan(&result.Id)
		
		if err == nil {
			return result, nil
		}else{
			return Client{}, errors.New("Unable to create Client in the database: "+err.Error())
		}
	}else{
		return result, errors.New("Unable to get a database connection to save the session")
	}
	
}

//func CreateClient(firstName string, lastName string, address string, city string, state string, zip string, phone string, entranceDate string, transactionDate string) (Client, error) {
//	
//	
//	
//}




