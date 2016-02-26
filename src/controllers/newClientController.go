package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"fmt"
)

type newClientController struct {
	template *template.Template
}

func (this *newClientController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vm := viewmodels.GetNewClient()
	fmt.Println(vm.Client)
	
	//This part handles what to do if the request was a post
//	if req.Method == "POST" {
//		vm.Client.SetName(req.FormValue("name"))
//		vm.Client.SetAddress(req.FormValue("address"))
//		vm.Client.SetPhone(req.FormValue("phone"))
//	}
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}

