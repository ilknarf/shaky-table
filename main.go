package main

import (
	"log"
	"net/http"

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

	log.Println("Initializing API router")
	apiRouter := api.NewRouter(userDB)
	http.Handle("/api", apiRouter)

	log.Println("Serving...")
	log.Fatal(http.ListenAndServeTLS(":8080", "ssl/host.crt", "ssl/host.key", nil))
}
