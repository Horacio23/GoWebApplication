package controllers

import (
	"GoWebApplication/src/controllers/util"
	"GoWebApplication/src/models"
	"GoWebApplication/src/viewmodels"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

const (
	timeFormat = "01/02/2006"
)

type clientsController struct {
	template *template.Template
}

func (this *clientsController) get(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetClients()
	clients, err := models.GetAllClients()

	if err == nil {
		vm.Clients = clients

		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		_, err := req.Cookie("sessionId")
		if err == nil {
			// get the member cookie and set the displayed name to the member
			if cookie, err := req.Cookie("user"); err == nil {

				vm.User = cookie.Value

			} else {
				fmt.Println("Error retrieving the member cookie:", err.Error())
			}

		} else {
			// if there is no session cookie then redirect to login
			http.Redirect(responseWriter, req, "/login", http.StatusFound)
		}

		this.template.Execute(responseWriter, vm)
	} else {
		fmt.Println("Error getting clients: " + err.Error())
		w.WriteHeader(404)
	}
}

func (this *clientsController) getClientsByTransaction(w http.ResponseWriter, req *http.Request) {
	vm := viewmodels.GetClients()

	vars := mux.Vars(req) //gets all the variables in the current request

	transaction := vars["transaction"]

	clients, err := models.GetClientsByTransaction(transaction)

	if err == nil {
		vm.Clients = clients

		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		_, err := req.Cookie("sessionId")
		if err == nil {
			// get the member cookie and set the displayed name to the member
			if cookie, err := req.Cookie("user"); err == nil {

				vm.User = cookie.Value

			} else {
				fmt.Println("Error retrieving the member cookie:", err.Error())
			}

		} else {
			// if there is no session cookie then redirect to login
			http.Redirect(responseWriter, req, "/login", http.StatusFound)
		}

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
	vars := mux.Vars(req) //gets all the variables in the current request

	idRaw := vars["id"]
	println("received Id " + idRaw)

	id, scErr := strconv.Atoi(idRaw) //Atoi stands for askii to int

	if scErr == nil {
		vm := viewmodels.GetClientView()
		w.Header().Add("Content-Type", "text/html")
		responseWriter := util.GetResponseWriter(w, req)
		defer responseWriter.Close()

		_, err := req.Cookie("sessionId")
		if err == nil {
			// get the member cookie and set the displayed name to the member
			if cookie, err := req.Cookie("user"); err == nil {

				vm.User = cookie.Value
				client, dbErr := models.GetClient(id)

				if dbErr == nil {

					vm.Client = client
					fmt.Println(client)

					this.template.Execute(responseWriter, vm)
				} else {
					fmt.Println("Error getting client: " + dbErr.Error())
					w.WriteHeader(404)
				}

			} else {
				fmt.Println("Error retrieving the member cookie:", err.Error())
			}

		} else {
			// if there is no session cookie then redirect to login
			http.Redirect(responseWriter, req, "/login", http.StatusFound)
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
				client.Email = req.FormValue("email")
				client.EntranceDate = req.FormValue("entranceDate")
				client.LastTransaction = req.FormValue("lastTransaction")
				client.TransactionDate = req.FormValue("transactionDate")
				client.Payment = sanitizePaymnet(req.FormValue("payment"))
				client.Notes = req.FormValue("notes")

				if client, ccErr := models.UpdateClient(client); ccErr == nil {
					vm.Client = client
					http.Redirect(responseWriter, req, "/clients", http.StatusFound)
				} else {
					fmt.Println("There was an error updating the client:", ccErr.Error())
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
	vm := viewmodels.GetNewClient()

	_, err := req.Cookie("sessionId")
	if err == nil {
		// get the member cookie and set the displayed name to the member
		if cookie, err := req.Cookie("user"); err == nil {

			vm.User = cookie.Value
			if req.Method == "POST" {
				println("POST method received for New Client")

				client.FirstName = req.FormValue("firstName")
				client.LastName = req.FormValue("lastName")
				client.Address = req.FormValue("address")
				client.City = req.FormValue("city")
				client.State = req.FormValue("state")
				client.Zip = req.FormValue("zip")
				client.Phone = req.FormValue("phone")
				client.Email = req.FormValue("email")
				client.EntranceDate = req.FormValue("entranceDate")
				client.LastTransaction = req.FormValue("lastTransaction")
				client.TransactionDate = req.FormValue("transactionDate")
				// if no payment was specified, then set payment to 0.00
				if client.Payment = sanitizePaymnet(req.FormValue("payment")); client.Payment == "" {
					client.Payment = "0.00"
				}
				client.Notes = req.FormValue("notes")

				fmt.Println("Client post:", client)
				_, err := models.CreateClient(client)

				if err == nil {
					fmt.Println("Client was successfully added")
					http.Redirect(responseWriter, req, "/clients", http.StatusFound)
				} else {
					fmt.Println("There was a problem creating the client:", err.Error())
					responseWriter.WriteHeader(404)
				}

			} else {

				responseWriter.Header().Add("Content-Type", "text/html")

				this.template.Execute(responseWriter, vm)
			}

		} else {
			fmt.Println("Error retrieving the member cookie:", err.Error())
		}

	} else {
		// if there is no session cookie then redirect to login
		http.Redirect(responseWriter, req, "/login", http.StatusFound)
	}
	//This part handles what to do if the request was a post

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

func sanitizePaymnet(payment string) string {
	// the payment field comes in the form of $123.123 this will ignore everything but the number essentials
	reg, err := regexp.Compile("[^0-9\\.\\-]")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(payment, "")
}

func getFiveYearClients(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	if clients, err := models.CheckDates(); err == nil {
		for _, v := range clients {
			if date, err := time.Parse(timeFormat, v.EntranceDate); err == nil {
				fmt.Println(date.Sub(time.Now()))
			} else {
				fmt.Println("Error parsing the date in getFiveYearClients:", err.Error())
			}

		}
	}
}
