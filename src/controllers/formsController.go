package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
	"controllers/util"
)

type formsController struct {
	template *template.Template
}

func (this *formsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetForms()
	
	w.Header().Add("Content-Type", "text/html")
	responseWriter  := util.GetResponseWriter(w , req)
	defer responseWriter.Close()
	this.template.Execute(responseWriter,vm)
}

