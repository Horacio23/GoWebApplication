package controllers

import (
	"controllers/util"
	"fmt"
	"net/http"
	"text/template"
	"viewmodels"
)

type formsController struct {
	template *template.Template
}

func (this *formsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetForms()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	_, err := req.Cookie("sessionId")
	if err == nil {
		// get the member cookie and set the displayed name to the member
		if cookie, err := req.Cookie("user"); err == nil {

			vm.User = cookie.Value

		} else {
			fmt.Println("Error retrieving the member cookie:", err.Error())
		}

	} else {
		// if there is no session cookie then redirect to login
		http.Redirect(responseWriter, req, "/login", http.StatusFound)
	}

	this.template.Execute(responseWriter, vm)
}
