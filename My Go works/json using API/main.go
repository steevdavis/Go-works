package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"os"
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


func AlJsonEndPoint(w http.ResponseWriter, r *http.Request) {
	
	config, _ :=LoadConfiguration("config.json")
	fmt.Fprintln(w,config)
}



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/view", AlJsonEndPoint).Methods("GET")
	
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
}
