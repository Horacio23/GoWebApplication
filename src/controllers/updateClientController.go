package controllers

import (
	"controllers/util"
	"net/http"
	"text/template"
	"viewmodels"
	"github.com/gorilla/mux"
	"strconv"
)

type updateClientController struct {
	template *template.Template
}

func (this *updateClientController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	
	vars := mux.Vars(req) 	//gets all the variable sin the current request
	
	idRaw := vars["id"]
	
	id, err := strconv.Atoi(idRaw)	//Atoi stands for askii to int
	
	if err ==nil {
		vm := viewmodels.GetUpdateClient(id)
		
//		//This part handles what to do if the request was a post
//		if req.Method == "POST" {
//			vm.Client.SetName(req.FormValue("name"))
//			vm.Client.SetAddress(req.FormValue("address"))
//			vm.Client.SetPhone(req.FormValue("phone"))
//		}
	
		w.Header().Add("Content-Type", "text/html")
		responseWriter  := util.GetResponseWriter(w , req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter,vm)
	}else{
		println(err.Error())
		w.WriteHeader(404)
	}
	
}

