package main

import(
	"log"
	"net/http"
	"html/template"
	"os"
	"server/pkg/file-watcher"
	"encoding/json"
	"math/rand/v2"
	"strconv"
	"io/ioutil"
	"strings"
)

type PageData struct {
	Title string
} 

type Cries struct {
		Latest string
		Legacy string
}

type PokemonData struct {
	Name string
	FrontDefault string
	Cries Cries 
}

type Pokemon struct {
	ID  int `json:"id"`
	Name  string `json:"name"`
	Sprites  struct {
		FrontDefault string `json:"front_default"`
	} `json:"sprites"`
	Cries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
}

var getPokimonByIdxUrl string = "https://pokeapi.co/api/v2/pokemon/"


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

	// generate random idx

	randPokeUrl := getPokimonByIdxUrl + strconv.Itoa(rand.IntN(152))
	log.Println(randPokeUrl)

	res, err := http.Get(randPokeUrl)

	if err != nil {
		log.Printf("Error getting pokemon: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading pokemon response: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var pokemonData Pokemon
	json.Unmarshal(responseData, &pokemonData)
	pokemon := PokemonData{
		strings.Title(pokemonData.Name), 
		pokemonData.Sprites.FrontDefault, 
		Cries{
			Latest:pokemonData.Cries.Latest, 
			Legacy:pokemonData.Cries.Legacy,
		},
	}

	dir, err := os.Getwd()
	if err != nil{
		log.Fatalf("Error getting current working directory: %v", err)
	}

	pathToIndexHtml := dir + "/public/pokemon.html"

	tmpl, err := template.ParseFiles(pathToIndexHtml)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, pokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
