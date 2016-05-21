package models

import "fmt"

type Transaction struct {
	ID          int
	ClientID    string
	Transaction string
	Amount      string
	Date        string
}

func getTransactions(query string) ([]Transaction, error) {
	result := []Transaction{}

	if db, dbErr := getDBConnection(); dbErr == nil {
		defer db.Close()

		if row, err := db.Query(query); err == nil {
			ts := Transaction{}
			for row.Next() {
				row.Scan(&ts.ID, &ts.ClientID, &ts.Transaction, &ts.Amount, &ts.Date)

				fmt.Println("The object is", ts)
				result = append(result, ts)
			}
		} else {
			return nil, err
		}

	} else {
		return nil, dbErr
	}

	return result, nil
}

func GetTransactionsByProcess(transaction string) ([]Transaction, error) {
	return getTransactions("select * from transactions where transaction='" + transaction + "'")
}

func GetAllTransactions() ([]Transaction, error) {
	return getTransactions("select * from transactions")
}
