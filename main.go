package main
import "log"
//main function
//loads config
func main(){
	LoadConfig()
	log.Printf(Config.Database.Dbname)
}