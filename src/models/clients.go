package models

import (
	"GoWebApplication/src/controllers/util"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/leekchan/accounting"
)

type Client struct {
	Id              int
	FirstName       string
	LastName        string
	Address         string
	City            string
	State           string
	Zip             string
	Phone           string
	Email           string
	LastTransaction string
	EntranceDate    string
	TransactionDate string
	Payment         string
	Notes           string
}

func GetAllClients() ([]Client, error) {
	return getClients(`SELECT * FROM clients`)
}

func GetClientsByTransaction(transaction string) ([]Client, error) {
	query := "SELECT * FROM clients WHERE last_transaction='" + transaction + "'"
	return getClients(query)
}

func getClients(query string) ([]Client, error) {
	//TODO: get list of clients from the database
	result := []Client{}

	db, dbErr := getDBConnection()

	if dbErr == nil {
		defer db.Close()

		row, qErr := db.Query(query)

		if qErr == nil {
			var id int
			var firstName string
			var lastName string
			var address string
			var city string
			var state string
			var zip string
			var phone string
			var email string
			var lastTransaction string
			var payment string
			var notes string
			var entranceDate time.Time
			var transactionDate time.Time

			defer row.Close()

			for row.Next() {
				sErr := row.Scan(&id, &firstName, &lastName, &address, &city, &state, &zip, &phone, &email, &entranceDate, &lastTransaction, &transactionDate, &payment, &notes)
				if sErr == nil {
					client := Client{
						Id:              id,
						FirstName:       firstName,
						LastName:        lastName,
						Address:         address,
						City:            city,
						State:           state,
						Zip:             zip,
						Phone:           phone,
						Email:           email,
						EntranceDate:    util.GetDate(entranceDate),
						LastTransaction: lastTransaction,
						TransactionDate: util.GetDate(transactionDate),
						Payment:         payment,
						Notes:           notes,
					}

					result = append(result, client)

				} else {
					return result, errors.New("Unable to get a field from the result query:" + sErr.Error())
				}
			}

			return result, nil

		} else {
			return result, errors.New("Unable to run query")
		}
	} else {
		return result, errors.New("Unable to get a database connection to save the session")
	}
}

func GetClient(id int) (Client, error) {
	result := Client{}

	var entranceDate time.Time
	var transactionDate time.Time
	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		dbErr := db.QueryRow(`SELECT id, first_name, last_name, address, city, state, zip, phone,
        email, entrance_date, last_transaction, transaction_date, payment, notes
  		FROM clients where id=$1`, id).Scan(&result.Id, &result.FirstName, &result.LastName, &result.Address, &result.City, &result.State, &result.Zip, &result.Phone, &result.Email, &entranceDate, &result.LastTransaction, &transactionDate, &result.Payment, &result.Notes)

		result.EntranceDate = util.GetDate(entranceDate)
		result.TransactionDate = util.GetDate(transactionDate)
		if payment, strErr := strconv.ParseFloat(result.Payment, 10); strErr != nil {
			fmt.Println("Error parsing payment", strErr.Error())
		} else {
			result.Payment = ac.FormatMoney(payment)
		}

		if dbErr == nil {
			fmt.Println("no error")
			fmt.Println(result)
			return result, nil
		} else {
			fmt.Println("error in query")
			return Client{}, errors.New("Unable to get Client from the database: " + dbErr.Error())
		}
	} else {
		fmt.Println("error with the db")
		return result, errors.New("Unable to get a database connection to save the session")
	}

}

func CreateClient(client Client) (Client, error) {
	fmt.Println("Inside Create Client")

	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		err := db.QueryRow(`INSERT INTO clients
			(first_name, last_name, address, city, state, zip, phone, email, entrance_date, last_transaction, transaction_date, payment, notes)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
			RETURNING id`, client.FirstName, client.LastName, client.Address, client.City, client.State, client.Zip, client.Phone, client.Email, client.EntranceDate, client.LastTransaction, client.TransactionDate, client.Payment, client.Notes).Scan(&client.Id)

		if err == nil {
			fmt.Println("The create query ran without errors")
			return client, nil
		} else {
			fmt.Println("there was a problem creating the client:", err.Error())
			return Client{}, errors.New("Unable to create Client in the database: " + err.Error())
		}
	} else {
		return client, errors.New("Unable to get a database connection to save the session")
	}

}

func UpdateClient(client Client) (Client, error) {

	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		_, err := db.Query(`UPDATE clients
			set first_name=$1, last_name=$2, address=$3, city=$4, state=$5, zip=$6, phone=$7, email=$8, entrance_date=$9, last_transaction=$10, transaction_date=$11, payment=$12, notes=$13
			WHERE id=$14`, client.FirstName, client.LastName, client.Address, client.City, client.State, client.Zip, client.Phone, client.Email, client.EntranceDate, client.LastTransaction, client.TransactionDate, client.Payment, client.Notes, client.Id)

		if err == nil {
			return client, nil
		} else {
			return Client{}, errors.New("Unable to create Client in the database: " + err.Error())
		}
	} else {
		return client, errors.New("Unable to get a database connection to save the session")
	}

}

func DeleteClient(id int) (bool, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		if _, dbErr := db.Query(`DELETE FROM clients WHERE id=$1`, id); dbErr == nil {
			return true, nil
		} else {
			return false, errors.New("Unable to delete the client: " + dbErr.Error())
		}

	} else {
		return false, errors.New("Unalbe to get a database connection in delete" + err.Error())
	}

}

func CheckDates() ([]Client, error) {
	db, err := getDBConnection()

	if err == nil {
		defer db.Close()
		if clients, dbErr := getClients(`SELECT * FROM clients WHERE entrance_date > now()::date-365`); dbErr == nil {
			return clients, nil
		} else {
			return nil, errors.New("Unable to run the check dates method: " + dbErr.Error())
		}

	} else {
		return nil, errors.New("Unable to get a database connection in checkDate" + err.Error())
	}

}
