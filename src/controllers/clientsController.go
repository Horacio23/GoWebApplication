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
				client.FirstName = req.FormValue("firstName")
				client.LastName = req.FormValue("lastName")
				client.Address = req.FormValue("address")
				client.City = req.FormValue("city")
				client.State = req.FormValue("state")
				client.Zip = req.FormValue("zip")
				client.Phone = req.FormValue("phone")
				client.EntranceDate = req.FormValue("entranceDate")
				client.LastTransaction = req.FormValue("lastTransaction")
				client.TransactionDate = req.FormValue("transactionDate")
				client.Notes = req.FormValue("notes")

				if client, ccErr := models.UpdateClient(client); ccErr == nil {
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

	client := models.Client{}

	//This part handles what to do if the request was a post
	if req.Method == "POST" {
		println("POST method received for New Client")

		client.FirstName = req.FormValue("firstName")
		client.LastName = req.FormValue("lastName")
		client.Address = req.FormValue("address")
		client.City = req.FormValue("city")
		client.State = req.FormValue("state")
		client.Zip = req.FormValue("zip")
		client.Phone = req.FormValue("phone")
		client.EntranceDate = req.FormValue("entranceDate")
		client.LastTransaction = req.FormValue("lastTransaction")
		client.TransactionDate = req.FormValue("transactionDate")
		client.Notes = req.FormValue("notes")

		_, err := models.CreateClient(client)

		if err == nil {
			fmt.Println("Client was successfully added")
			http.Redirect(responseWriter, req, "/clients", http.StatusFound)
		} else {
			fmt.Println("There was a problem creating the client:", err.Error())
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
