package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
)

type newClientController struct {
	template *template.Template
}

func (this *newClientController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetNewClient()
	
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

