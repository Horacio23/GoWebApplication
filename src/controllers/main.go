package controllers

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
)

func Register(templates *template.Template) {
	println("REGISTERING TEMPLATE")

	router := mux.NewRouter()

	homeController := new(homeController)
	homeController.template = templates.Lookup("index.html")
	homeController.loginTemplate = templates.Lookup("login.html")
	homeController.signupTemplate = templates.Lookup("signup.html")
	router.HandleFunc("/home", homeController.get)
	router.HandleFunc("/login", homeController.login)
	router.HandleFunc("/signup", homeController.signup)
	router.HandleFunc("/logout", homeController.logout)

	formsController := new(formsController)
	formsController.template = templates.Lookup("forms.html")
	router.HandleFunc("/forms", formsController.get)

	clientsController := new(clientsController)
	clientsController.template = templates.Lookup("clients.html")
	router.HandleFunc("/clients", clientsController.get)
	router.HandleFunc("/clients/{transaction}", clientsController.getClientsByTransaction)

	cController := new(clientController)
	cController.template = templates.Lookup("clientModal.html")
	router.HandleFunc("/client/{id}", cController.get)

	newClientController := new(clientController)
	newClientController.template = templates.Lookup("newClient.html")
	router.HandleFunc("/newClient", newClientController.post)

	updateClientController := new(clientController)
	updateClientController.template = templates.Lookup("newClient.html") //same page as new client
	router.HandleFunc("/update/{id}", updateClientController.update)

	deleteClientController := new(clientController)
	router.HandleFunc("/delete/{id}", deleteClientController.remove)

	http.Handle("/", router)

	//Resources
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/font-awesome/", serveResource)
	http.HandleFunc("/fonts/", serveResource)
	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/bootstrap-datepicker/", serveResource)
}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "public" + req.URL.Path

	var contentType string

	if strings.Contains(path, ".css") {
		contentType = "text/css"
	} else if strings.Contains(path, ".png") {
		contentType = "image/png"
	} else if strings.Contains(path, ".js") {
		contentType = "text/javascript"
	} else {
		contentType = "text/plain"
	}

	f, err := os.Open(path)

	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)

		br := bufio.NewReader(f) //buffered readers are important to not load the full page in one go
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
