package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Singleton config
var Config Configuration
//struct defining config yml file data structure
type Configuration struct {
	Database database
}

type database struct {
	Host     string
	User     string
	Password string
	Port     int
	Dbname   string
	Sslmode  string
}
//func to load configuration from file to Configuration struct
func LoadConfig() {
	source, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &Config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", Config.Database)
}
