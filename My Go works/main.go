package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	TraceID   string `json:"trace_id"`
	Timestamp int64  `json:"timestamp"`
	Result    string `json:"result"`
	Metadata  struct {
		Service struct {
			Name string `json:"name"`
		} `json:"service"`
		Version  interface{} `json:"version"`
		Language struct {
			Version string `json:"version"`
			Name    string `json:"name"`
		} `json:"language"`
		Agent struct {
			Version string `json:"version"`
			Name    string `json:"name"`
		} `json:"agent"`
		Framework struct {
			Version string `json:"version"`
			Name    string `json:"name"`
		} `json:"framework"`
	} `json:"metadata"`
	Request struct {
		URL     string `json:"url"`
		Body    string `json:"body"`
		Headers struct {
			ContentLength string `json:"content-length"`
			ContentType   string `json:"content-type"`
			Host          string `json:"host"`
			Accept        string `json:"accept"`
			UserAgent     string `json:"user-agent"`
		} `json:"headers"`
		Method string `json:"method"`
	} `json:"request"`
	Response struct {
		StatusCode int `json:"status_code"`
		Headers    struct {
			ContentLength string `json:"Content-Length"`
			ContentType   string `json:"Content-Type"`
		} `json:"headers"`
	} `json:"response"`
	Duration float64 `json:"duration"`
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
	
