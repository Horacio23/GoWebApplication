package models

import "fmt"

type Transaction struct {
	ID          int
	ClientID    string
	Transaction string
	Amount      string
	Date        string
}

func GetTransactions() ([]Transaction, error) {
	result := []Transaction{}

	if db, dbErr := getDBConnection(); dbErr == nil {
		defer db.Close()

		if row, err := db.Query("select * from transactions"); err == nil {
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
