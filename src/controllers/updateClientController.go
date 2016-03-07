package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
	"models"
)

type updateClientController struct {
	template *template.Template
}

func (this *updateClientController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vars := mux.Vars(req) 	//gets all the variable sin the current request
	
	idRaw := vars["id"]
	
	id, scErr := strconv.Atoi(idRaw)	//Atoi stands for askii to int
	
	if scErr ==nil {
		vm := viewmodels.GetUpdateClient()
		
		if client, err := models.GetClient(id); err == nil{
			//this is the client that is to be updated
			vm.Client = client
			
			//This part handles what to do if the request was a post
			if req.Method == "POST" {
				firstName := req.FormValue("firstName")
				lastName := req.FormValue("lastName")
				address := req.FormValue("address")
				city := req.FormValue("city")
				state := req.FormValue("state")
				zip := req.FormValue("zip")
				phone := req.FormValue("phone")
				entranceDate := req.FormValue("entranceDate")
				transactionDate := req.FormValue("transactionDate")
				
		
				if client, ccErr := models.UpdateClient(id, firstName, lastName, address, city, state, zip, phone, entranceDate, transactionDate); ccErr == nil{
					vm.Client = client
					http.Redirect(responseWriter, req, "/clients", http.StatusFound)
				}
			}
		
			w.Header().Add("Content-Type", "text/html")
			responseWriter  := util.GetResponseWriter(w , req)
			defer responseWriter.Close()
			this.template.Execute(responseWriter,vm)
		}
		
	
	}else{
		println(scErr.Error())
		w.WriteHeader(404)
	}
	
}

