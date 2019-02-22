package main

import (
	"encoding/json"
	"fmt"
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


func main() {
	fmt.Println("start")
	config, _ :=LoadConfiguration("config.json")
	fmt.Println(config)
}
	
