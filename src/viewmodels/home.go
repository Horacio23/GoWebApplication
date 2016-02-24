package viewmodels

import (
	
)

type Home struct {
	Title string
	Active string
	Id int
}

func GetHome() Home {
	result := Home{
		Title: "Home",
		Active: "home",
		Id:1,
	}
	
	return result
}