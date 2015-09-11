package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Singleton config
var Config Configuration

type Configuration struct {
	Database Database
}

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Dbname   string
	Sslmode  string
}

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
