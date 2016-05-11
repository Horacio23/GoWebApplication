package viewmodels

type Forms struct {
	Title  string
	Active string
	User   string
	Forms  []Form
}

type Form struct {
	Title   string
	Content string
	User    string
	Id      int
}

func GetForms() Forms {

	result := Forms{
		Title:  "Forms",
		Active: "forms",
	}

	form1 := Form{
		Title:   "I-130",
		Content: "The content for the i-130 form",
		Id:      1,
	}

	form2 := Form{
		Title:   "1099",
		Content: "The other conntent",
		Id:      2,
	}

	form3 := Form{
		Title:   "From3",
		Content: "The conten form3",
		Id:      3,
	}

	result.Forms = []Form{
		form1,
		form2,
		form3,
	}

	return result

}
