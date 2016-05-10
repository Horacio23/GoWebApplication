package controllers

import (
	"GoWebApplication/src/controllers/util"
	"GoWebApplication/src/viewmodels"
	"net/http"
	"text/template"
)

type formsController struct {
	template *template.Template
}

func (this *formsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetForms()

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter, vm)
}
