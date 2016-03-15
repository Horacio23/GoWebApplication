package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"models"
)

type newClientController struct {
	template *template.Template
}

func (this *newClientController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	
	
	
	
	//This part handles what to do if the request was a post
	if req.Method == "POST" {
		println("POST method received for New Client")
		
		vm := viewmodels.GetClientView()
		
		firstName := req.FormValue("firstName")
		lastName := req.FormValue("lastName")
		address := req.FormValue("address")
		city := req.FormValue("city")
		state := req.FormValue("state")
		zip := req.FormValue("zip")
		phone := req.FormValue("phone")
		entranceDate := req.FormValue("entranceDate")
		lastTransaction := req.FormValue("lastTransaction")
		transactionDate := req.FormValue("transactionDate")
		notes := req.FormValue("notes")
		
		if(firstName!="" && lastName!="" && address!="" && city!="" && state!="" && entranceDate!="" && transactionDate!=""){
			client, err := models.CreateClient(firstName, lastName, address, city, state, zip, phone, entranceDate, lastTransaction, transactionDate, notes)
			
			if err == nil{
	
				vm.Client = client
				http.Redirect(responseWriter, req, "/clients", http.StatusFound)
			}
		}else{
			responseWriter.WriteHeader(404)
		}
		
	}else{
		vm := viewmodels.GetNewClient()
		responseWriter.Header().Add("Content-Type", "text/html")
		this.template.Execute(responseWriter, vm)
	}
}

