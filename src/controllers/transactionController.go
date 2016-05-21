package controllers

import (
	"GoWebApplication/src/models"
	"GolangWebApp/src/controllers/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllTransactions(w http.ResponseWriter, req *http.Request) {

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "application/json")

	if trans, err := models.GetTransactions(); err == nil {
		if response, err := json.Marshal(trans); err == nil {
			fmt.Println(response)
			json.NewEncoder(w).Encode(response)
		} else {
			fmt.Println("Error creating a json from the transactions", err.Error())
		}
	} else {
		fmt.Println("There was an error getting the transactions", err.Error())
		w.WriteHeader(404)
	}

}
