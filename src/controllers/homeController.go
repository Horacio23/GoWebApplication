package controllers

import (
	"GoWebApplication/src/controllers/util"
	"GoWebApplication/src/models"
	"GoWebApplication/src/viewmodels"
	"fmt"
	"net/http"
	"text/template"
	"time"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	println("Getting Home controller")

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

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetLogin()
	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")

		member, err := models.GetMember(username, password)

		if err == nil {
			session, err := models.CreateSession(member)
			if err == nil {
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId
				cookie.Expires = time.Now().Add(1 * time.Hour)
				responseWriter.Header().Add("Set-Cookie", cookie.String())

				var cookieM http.Cookie
				cookieM.Name = "user"
				cookieM.Value = member.FirstName
				cookieM.Expires = time.Now().Add(1 * time.Hour)
				responseWriter.Header().Add("Set-Cookie", cookie.String())

			}

			vmh := viewmodels.GetHome()

			vmh.User = member.FirstName

			this.template.Execute(responseWriter, vmh)
		} else {
			fmt.Println("Unable to get member in Login:", err.Error())
			vm.Error = "Invalid Username or Password. Please try again"
		}

	}

	this.loginTemplate.Execute(responseWriter, vm)
}
