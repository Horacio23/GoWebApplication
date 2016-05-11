package viewmodels

type Home struct {
	Title  string
	Active string
	User   string
	Id     int
}

func GetHome() Home {
	result := Home{
		Title:  "Home",
		Active: "home",
		Id:     1,
	}

	return result
}

type Login struct {
	Title  string
	Active string
	User   string
	Error  string
}

func GetLogin() Login {
	result := Login{
		Title:  "Immigration Program - Login",
		Active: "",
	}

	return result
}
