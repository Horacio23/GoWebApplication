package controllers

import (
	"net/http"
	"text/template"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
	"controllers/util"
)


type clientController struct {
	template *template.Template
}

func (this *clientController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req) 	//gets all the variable sin the current request
	
	idRaw := vars["id"]
	
	id, err := strconv.Atoi(idRaw)	//Atoi stands for askii to int
	
	if err ==nil {
		vm := viewmodels.GetClient(id)
	
		w.Header().Add("Content-Type", "text/html")
		responseWriter  := util.GetResponseWriter(w , req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter,vm)
	}else{
		w.WriteHeader(404)
	}
}