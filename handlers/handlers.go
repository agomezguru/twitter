package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/agomezguru/twitter/middlewares"
	"github.com/agomezguru/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Start Web Server in selected PORT */ 
func Drivers() {
	router := mux.NewRouter()
	
	// All functional endpoints go here
	router.HandleFunc("/register", middlewares.DBVerify(routers.Register)).Methods("POST")
	// try to find env setting 
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // default OS port
	}
	
	// Set access permits to the world 
	handler := cors.AllowAll().Handler(router)

	// Init web server
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}