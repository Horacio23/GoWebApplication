package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
	"controllers/util"
	"models"
)

type homeController struct {
	template *template.Template
	loginTemplate *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetHome()
	println("Getting Home controller")
	
	w.Header().Add("Content-Type", "text/html")
	responseWriter  := util.GetResponseWriter(w , req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter,vm)
	
}

func (this *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	responseWriter.Header().Add("Content-Type", "text/html")
	
	if req.Method == "POST" {
		username := req.FormValue("username")
		password := req.FormValue("password")
		
		member, err := models.GetMember(username, password)
		
		if err == nil{
			session, err :=models.CreateSession(member)
			if err == nil{
				var cookie http.Cookie
				cookie.Name = "sessionId"
				cookie.Value = session.SessionId
				responseWriter.Header().Add("Set-Cookie", cookie.String())
				
			}
		}
	}
	
	vm := viewmodels.GetLogin()
	
	this.loginTemplate.Execute(responseWriter, vm)
}