package main

import(
	"log"
	"net/http"
	"html/template"
	"os"
)

type PageData struct {
	Title string
	Message string
} 


func main() {
	http.HandleFunc("/", handler)

	log.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed to load %v", err)
	}

}

func handler(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title: "Go Template with HTMX",
		Message: "Hello from the server",
	}

	dir, err := os.Getwd()
	if err != nil{
		log.Fatalf("Error getting current working directory: %v", err)
	}

	pathToIndexHtml := dir + "/public/index.html"

	log.Println("Current working directory:", pathToIndexHtml)

	tmpl, err := template.ParseFiles(pathToIndexHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(os.Stdout,pathToIndexHtml, data); err != nil {
		log.Printf("Error executing template: %v", err)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}