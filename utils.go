package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
}

var config Configuration

//P Convenience function for printing to stdout
func p(a ...interface{}) {
	fmt.Println(a...)
}

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatalln("Cannot open config file: ", err)
	}

	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Cannot get configuration from file", err)
	}
}

// Version number
func version() string {
	return "0.1"
}
