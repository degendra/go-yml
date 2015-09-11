package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Config is a singleton configuration holder
var Config Configuration

//Configuration holds config data
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
//LoadConfig parses file to Configuration struct
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
