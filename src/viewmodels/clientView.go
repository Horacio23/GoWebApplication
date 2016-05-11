package viewmodels

import (
	"GoWebApplication/src/models"
)

type ClientView struct {
	Title  string
	Active string
	User   string
	Client models.Client
}

func GetClientView() ClientView {
	result := ClientView{
		Title:  "Client",
		Active: "clients",
	}

	return result

}
