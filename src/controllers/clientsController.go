package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
	"controllers/util"
	"models"
	"fmt"
)

type clientsController struct {
	template *template.Template
}

func (this *clientsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetClients()
	clients, err := models.GetClients()
	
	if err == nil {
		vm.Clients = clients
	
		w.Header().Add("Content-Type", "text/html")
		responseWriter  := util.GetResponseWriter(w , req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter,vm)
	}else{
		fmt.Println("Error getting clients: "+err.Error())
		w.WriteHeader(404)
	}
}

type clientController struct {
	template *template.Template
}

func (this *clientController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) 	//gets all the variable sin the current request
	
	idRaw := vars["id"]
	println("received Id "+idRaw)
	
	id, scErr := strconv.Atoi(idRaw)	//Atoi stands for askii to int
	
	if scErr == nil {
		vm := viewmodels.GetClientView()
	
		client, dbErr := models.GetClient(id)
		
		if dbErr ==nil {
			
			vm.Client = client
			fmt.Println(client)
			w.Header().Add("Content-Type", "text/html")
			responseWriter  := util.GetResponseWriter(w , req)
			defer responseWriter.Close()
			this.template.Execute(responseWriter,vm)
		}else{
			fmt.Println("Error getting client: "+dbErr.Error())
			w.WriteHeader(404)
		}
	}else{
		fmt.Println("Error converting string to int: "+scErr.Error())
	}
	
}