package main

import (
	"fmt"
	"encoding/json"
	"log"
	//"math/rand"
	"net/http"
	//"strconv"
        "os"
	"github.com/gorilla/mux"
)

type Config struct {
	
	Metadata  struct {
		Service struct {
			Name string `json:"name"`
		} `json:"service"`
		Version  interface{} `json:"version"`
		Language struct {
			Version string `json:"version"`
			Name    string `json:"name"`
		} `json:"language"`
		
	} `json:"metadata"`
	
	
}


func LoadConfiguration(filename string) (Config, error) {
    var config Config
    configFile, err := os.Open(filename)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config, err
}
  

func getJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	//json.NewEncoder(w).Encode(books)
	
	fmt.Println("start")
	config, _ :=LoadConfiguration("config.json")
	fmt.Println(config)
}




func main() {

        r := mux.NewRouter()
r.HandleFunc("/jsn", getJson).Methods("GET")
log.Fatal(http.ListenAndServe(":8000", r))	
}
	
