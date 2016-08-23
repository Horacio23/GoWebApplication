package main

import (
	"controllers"
	"log"
	"net/http"
	"os"
	"text/template"
)

func main() {
	//	http.HandleFunc("/", func( w http.ResponseWriter, req *http.Request) {
	//			w.Header().Add("Content Type", "text/html")
	//			templates := template.New("template")
	//			templates.New("test").Parse(doc)
	//			templates.New("header").Parse(header)
	//			templates.New("footer").Parse(footer)
	//
	//
	//			context := Context{
	//				[3]string{"Lemon", "Orange", "Apple"},
	//				"the title",
	//			}
	//			templates.Lookup("test").Execute(w, context)
	//
	//		})

	templates := populateTemplates()

	controllers.Register(templates)

	// port := os.Getenv("PORT")
	//
	// if port == "" {
	// 	log.Fatal("$PORT must be set")
	// }

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		log.Fatal("The server went to shit:", err)
	}
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close() //this line will execute after the func is done

	templatePathsRaw, _ := templateFolder.Readdir(-1) //this func returns a specific number of file paths every time its called, -1 returns all paths

	templatePaths := new([]string)

	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())

		}
	}

	result.ParseFiles(*templatePaths...) //allows the compiler to break the array into multiple input

	println("TMEPLATES POPULATED")
	return result
}
