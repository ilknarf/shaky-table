package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ilknarf/shaky-table/api"
	"github.com/ilknarf/shaky-table/userdb"
)

func main() {
	log.Println("Starting...")

	log.Println("Connecting to UserDB")
	userDB, err := userdb.Open()

	if err != nil {
		log.Fatalln(err)
	}
	defer userDB.Close()

	log.Println("Initializing router")
	baseRouter := mux.NewRouter()

	// add apiRouter to baseRouter
	log.Println("Adding API routes to router")
	apiRouter := api.NewRouter(userDB)
	baseRouter.PathPrefix("/api").Name("api").Subrouter().Handle("/", apiRouter)

	log.Println("Serving...")
	log.Fatal(http.ListenAndServeTLS(":8080", "ssl/host.crt", "ssl/host.key", baseRouter))
}
