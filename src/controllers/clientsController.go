package controllers

import (
	"GoWebApplication/src/controllers/util"
	"GoWebApplication/src/models"
	"GoWebApplication/src/viewmodels"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
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
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	} else {
		fmt.Println("Error getting clients: " + err.Error())
		w.WriteHeader(404)
	}
}

type clientController struct {
	template *template.Template
}

func (this *clientController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) //gets all the variable sin the current request

	idRaw := vars["id"]
	println("received Id " + idRaw)

	id, scErr := strconv.Atoi(idRaw) //Atoi stands for askii to int

	if scErr == nil {
		vm := viewmodels.GetClientView()

		client, dbErr := models.GetClient(id)

		if dbErr == nil {

			vm.Client = client
			fmt.Println(client)
			w.Header().Add("Content-Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()
			this.template.Execute(responseWriter, vm)
		} else {
			fmt.Println("Error getting client: " + dbErr.Error())
			w.WriteHeader(404)
		}
	} else {
		fmt.Println("Error converting string to int: " + scErr.Error())
	}

}

func (this *clientController) update(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req) //gets all the variable sin the current request

	idRaw := vars["id"]

	id, scErr := strconv.Atoi(idRaw) //Atoi stands for askii to int

	if scErr == nil {
		vm := viewmodels.GetUpdateClient()

		if client, err := models.GetClient(id); err == nil {
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
				lastTransaction := req.FormValue("lastTransaction")
				transactionDate := req.FormValue("transactionDate")
				notes := req.FormValue("notes")

				if client, ccErr := models.UpdateClient(id, firstName, lastName, address, city, state, zip, phone, entranceDate, lastTransaction, transactionDate, notes); ccErr == nil {
					vm.Client = client
					http.Redirect(responseWriter, req, "/clients", http.StatusFound)
				}
			}

			w.Header().Add("Content-Type", "text/html")
			responseWriter := util.GetResponseWriter(w, req)
			defer responseWriter.Close()
			this.template.Execute(responseWriter, vm)
		}

	} else {
		println(scErr.Error())
		w.WriteHeader(404)
	}

}

func (this *clientController) post(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	//This part handles what to do if the request was a post
	if req.Method == "POST" {
		println("POST method received for New Client")

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

		if firstName != "" && lastName != "" && address != "" && city != "" && state != "" && entranceDate != "" && transactionDate != "" {
			_, err := models.CreateClient(firstName, lastName, address, city, state, zip, phone, entranceDate, lastTransaction, transactionDate, notes)

			if err == nil {

				http.Redirect(responseWriter, req, "/clients", http.StatusFound)
			}
		} else {
			responseWriter.WriteHeader(404)
		}

	} else {
		vm := viewmodels.GetNewClient()
		responseWriter.Header().Add("Content-Type", "text/html")
		this.template.Execute(responseWriter, vm)
	}
}

func (this *clientController) remove(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req) //gets all the variable sin the current request

	idRaw := vars["id"]

	id, scErr := strconv.Atoi(idRaw) //Atoi stands for askii to int

	if scErr == nil {
		if _, err := models.DeleteClient(id); err == nil {
			http.Redirect(responseWriter, req, "/clients", http.StatusFound)
		} else {
			w.WriteHeader(404)
		}
	} else {
		w.WriteHeader(404)
	}
}
