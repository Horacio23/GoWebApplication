package controllers

import (
	"net/http"
	"os"
	"text/template"
	"bufio"
	"strings"
	"github.com/gorilla/mux"
)

func Register(templates *template.Template) {
	println("REGISTERING TEMPLATE")
	
	router := mux.NewRouter()
	
	hc := new(homeController)
	hc.template = templates.Lookup("index.html")
	router.HandleFunc("/home",hc.get)
	
	fc := new(formsController)
	fc.template = templates.Lookup("forms.html")
	router.HandleFunc("/forms",fc.get)

	cc := new(clientController)
	cc.template = templates.Lookup("clients.html")
	router.HandleFunc("/client/{id}",cc.get)
	
	http.Handle("/", router)
	
	//Resources
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/font-awesome/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
	http.HandleFunc("/js/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path
	
	var contentType string
	
	if strings.Contains(path, ".css") {
		contentType = "text/css"
	}else if strings.Contains(path, ".png") {
		contentType = "image/png"
	}else if strings.Contains(path, ".js") {
		contentType = "text/javascript"
	}else{
		contentType = "text/plain"
	}
	
	f, err := os.Open(path)
	
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		
		br := bufio.NewReader(f)  	//buffered readers are important to not load the full page in one go
		br.WriteTo(w)
	}else{
		w.WriteHeader(404)
	}
}
