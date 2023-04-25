package main

import (
	"fmt"
	"log"
	"os"

	"github.com/agomezguru/twitter/db"
	"github.com/agomezguru/twitter/handlers"
)

func main() {
	if db.DBConnectionAlive() == false {
		log.Fatal("DB connection failed")
		return 
	}
	fmt.Println("Hola mundo")
	fmt.Println(os.Getenv("MONGODB_PASS"))
	handlers.Drivers()
}