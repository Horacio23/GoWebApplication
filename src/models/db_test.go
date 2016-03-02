package models

import (
	"testing"
)

func Test_DBWorks(t *testing.T) {
	_, err := getDBConnection() 
	
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}