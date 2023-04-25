package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN is the DB object connection */
var MongoCN = DBConnect()
var clientOptions = options.Client().ApplyURI(
	"mongodb+srv://alagunas-s:" + os.Getenv("MONGODB_PASS") +
	"@gocluster.tlu3j90.mongodb.net/?retryWrites=true&w=majority")

/* DBConnect make a remote DB connection */
func DBConnect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("DB connection successfull.")
	return client
	
}

/* DBConnectionAlive ping the DB, if is reachable return true. */
func DBConnectionAlive() bool {
	err := MongoCN.Ping (context.TODO(), nil)
	if err != nil {
		fmt.Println("Error en ping: " + err.Error())
		fmt.Println("Revise que tenga conexión a internet y su contraseña de acceso.")
		return false
	}
	fmt.Println("Ping Ok")
	return true
}