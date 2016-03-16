package models

import (
	"errors"
	"fmt"
	"time"
	"controllers/util"
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
	LastTransaction string
	EntranceDate string
	TransactionDate string
	Notes string
}



func GetClients() ([]Client, error) {
	//TODO: get list of clients from the database
	result := []Client{}
	
	db, dbErr := getDBConnection()
	
	if dbErr == nil{
		defer db.Close()
		
		row, qErr := db.Query(`SELECT * FROM clients`)
		
		if qErr == nil{
			var id int
			var firstName string
			var lastName string
			var address string
			var city string
			var state string
			var zip string
			var phone string
			var lastTransaction string
			var notes string
			var entranceDate time.Time
			var transactionDate time.Time
			
			defer row.Close()
			
			for row.Next() {
				sErr := row.Scan(&id, &firstName, &lastName, &address, &city, &state, &zip, &phone, &entranceDate, &lastTransaction, &transactionDate, &notes)
				if sErr == nil {
					client := Client{
						Id: id,
						FirstName: firstName,
						LastName: lastName,
						Address: address,
						City: city,
						State: state,
						Zip: zip,
						Phone: phone,
						EntranceDate: util.GetDate(entranceDate),
						LastTransaction: lastTransaction,
						TransactionDate: util.GetDate(transactionDate),
						Notes: notes,
					}
					
					result = append(result, client)
					
				}else{
					return result, errors.New("Unable to get a field from the result query:" + sErr.Error())
				}
			}
			
			return result, nil
			
		}else{
			return result, errors.New("Unable to run query")
		}
	}else{
		return result, errors.New("Unable to get a database connection to save the session")
	}
}

func GetClient(id int) (Client, error) {
	result := Client{}
	
	var entranceDate time.Time 
	var transactionDate time.Time
	
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		dbErr := db.QueryRow(`SELECT id, first_name, last_name, address, city, state, zip, phone, 
       	entrance_date, last_transaction, transaction_date, notes
  		FROM clients where id=$1`,id).Scan(&result.Id, &result.FirstName, &result.LastName, &result.Address, &result.City, &result.State, &result.Zip, &result.Phone, &entranceDate, &result.LastTransaction, &transactionDate, &result.Notes)
		
		result.EntranceDate = util.GetDate(entranceDate)
		result.TransactionDate = util.GetDate(transactionDate)
		
		
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

func CreateClient(firstName string, lastName string, address string, city string, state string, zip string, phone string, entranceDate string,lastTransaction string, transactionDate string, notes string) (Client, error) {
	fmt.Println("Inside Create Client")
	result := Client{}
	result.FirstName = firstName
	result.LastName = lastName
	result.Address = address
	result.City = city
	result.Zip = zip
	result.State = state
	result.Phone = phone
	result.LastTransaction = lastTransaction
	result.EntranceDate = entranceDate
	result.TransactionDate = transactionDate
	result.Notes = notes
	
	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO clients
			(first_name, last_name, address, city, state, zip, phone, entrance_date, last_transaction, transaction_date, notes)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			RETURNING id`, firstName, lastName, address, city, state, zip, phone, entranceDate, lastTransaction ,transactionDate, notes).Scan(&result.Id)
		
		if err == nil {
			fmt.Println("The create query ran without errors")
			return result, nil
		}else{
			fmt.Println("there was a problem creating the client:",err.Error())
			return Client{}, errors.New("Unable to create Client in the database: "+err.Error())
		}
	}else{
		return result, errors.New("Unable to get a database connection to save the session")
	}
	
}

func UpdateClient(id int, firstName string, lastName string, address string, city string, state string, zip string, phone string, entranceDate string,lastTransaction string, transactionDate string, notes string) (Client, error) {
	result := Client{}
	result.Id = id
	result.FirstName = firstName
	result.LastName = lastName
	result.Address = address
	result.City = city
	result.Zip = zip
	result.State = state
	result.Phone = phone
	result.LastTransaction = lastTransaction
	result.EntranceDate = entranceDate
	result.TransactionDate = transactionDate
	result.Notes = notes
	
	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		_, err := db.Query(`UPDATE clients
			set first_name=$1, last_name=$2, address=$3, city=$4, state=$5, zip=$6, phone=$7, entrance_date=$8, last_transaction=$9, transaction_date=$10, notes=$11
			WHERE id=$12`, firstName, lastName, address, city, state, zip, phone, entranceDate, lastTransaction, transactionDate, notes, id)
		
		if err == nil {
			return result, nil
		}else{
			return Client{}, errors.New("Unable to create Client in the database: "+err.Error())
		}
	}else{
		return result, errors.New("Unable to get a database connection to save the session")
	}
	
}

func DeleteClient(id int) (bool, error) {
	db, err := getDBConnection()
	
	if err == nil {
		defer db.Close()
		if _,dbErr := db.Query(`DELETE FROM clients WHERE id=$1`, id); dbErr==nil {
			return true, nil
		}else{
			return false, errors.New("Unalbe to delete the client: "+dbErr.Error())
		} 

	}else{
		return false, errors.New("Unalbe to get a database connection in delete"+err.Error())
	}

}



