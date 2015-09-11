package main
import "log"

func main(){
	LoadConfig()
	log.Printf(Config.Database.Dbname)
}