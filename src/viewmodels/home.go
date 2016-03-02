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

type Login struct {
	Title string
	Active string
}

func GetLogin() Login {
	result := Login{
		Title: "Lemonade Stand Society - Login",
		Active: "",
	}
	
	return result
}