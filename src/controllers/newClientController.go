package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"models"
	"strconv"
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
		transactionDate := req.FormValue("transactionDate")
		client, err := models.CreateClient(firstName, lastName, address, city, state, zip, phone, entranceDate, transactionDate)
		
		if err == nil{
//			vm.Client.FirstName = firstName
//			vm.Client.LastName = lastName
//			vm.Client.Address = address
//			vm.Client.City = city
//			vm.Client.State = state
//			vm.Client.Zip = zip
//			vm.Client.Phone = phone
//			vm.Client.EntranceDate = entranceDate
//			vm.Client.TransactionDate = transactionDate

			vm.Client = client
			http.Redirect(responseWriter, req, "/client/"+strconv.Itoa(client.Id), http.StatusFound)
		}
		
	}else{
		vm := viewmodels.GetNewClient()
		responseWriter.Header().Add("Content-Type", "text/html")
		this.template.Execute(responseWriter, vm)
	}
}

