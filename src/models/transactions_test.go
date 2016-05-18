package models

import (
	"fmt"
	"testing"
)

func Test_GetAllTransactions(t *testing.T) {
	trans, err := GetTransactions()

	if err == nil {

		fmt.Println("The size of the array is", len(trans))
	} else {
		t.Log("Test for creating new client failed: " + err.Error())

		t.FailNow()
	}
}
