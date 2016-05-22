package models

import (
	"fmt"
	"testing"
)

func Test_GetAllTransactions(t *testing.T) {
	trans, err := GetAllTransactions()

	if err == nil {

		fmt.Println("The size of the transaction array is", len(trans))
	} else {
		t.Log("Test for creating new client failed: " + err.Error())

		t.FailNow()
	}
}
