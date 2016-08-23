package controllers

import (
	"controllers/util"
	"encoding/json"
	"fmt"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllTransactions(w http.ResponseWriter, req *http.Request) {

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "application/json")

	if trans, err := models.GetAllTransactions(); err == nil {
		if response, err := json.Marshal(trans); err == nil {

			responseWriter.Write(response)

		} else {
			fmt.Println("Method: getAllTransactions. Error creating a json from the transactions", err.Error())
		}
	} else {
		fmt.Println("There was an error getting the transactions", err.Error())
		w.WriteHeader(404)
	}

}

func getTransactionByProcess(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vars := mux.Vars(req) //gets all the variables in the current request

	process := vars["process"]

	responseWriter.Header().Add("Content-Type", "application/json")

	if trans, err := models.GetTransactionsByProcess(process); err == nil {
		if response, err := json.Marshal(trans); err == nil {

			responseWriter.Write(response)

		} else {
			fmt.Println("Method: getTransactionByProcess. Error creating a json from the transactions", err.Error())
		}
	} else {
		fmt.Println("There was an error getting the transactions", err.Error())
		w.WriteHeader(404)
	}
}
