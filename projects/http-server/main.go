package main

import(
	"log"
	"net/http"
	"html/template"
	"os"
	"server/pkg/file-watcher"
	"encoding/json"
)

type PageData struct {
	Title string
} 

type PokemonData struct {
	PokeName string
	PokeHp string
}


func main() {
	// start file watcher for development 
	go filewatcher.WatchFiles()

	http.HandleFunc("/", handler)
	http.HandleFunc("/get-pokemon", getPokemonHandler)

	log.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalf("Server failed to load %v", err)
	}

}

func getPokemonHandler(w http.ResponseWriter, r *http.Request){
	pokemon := PokemonData()

	res := map[string]interface{}{
		"PokeName" : "Pickachu",
		"PokeHp" : "50",
	}

	pokemon.PokeHp = res["PokeHp"].(string)
	pokemon.PokeName = res["PokeName"].(string)

	jsonResponse, err := json.Marshal(pokemon)
	if err != nil{
		log.Printf("Error marshaling JSON: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")

	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request){

	data := PageData{
		Title: "Go Template with HTMX",
	}

	dir, err := os.Getwd()
	if err != nil{
		log.Fatalf("Error getting current working directory: %v", err)
	}

	pathToIndexHtml := dir + "/public/index.html"

	tmpl, err := template.ParseFiles(pathToIndexHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
