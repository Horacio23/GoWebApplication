package controllers

import (
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
	this.template.Execute(w,vm)
}

