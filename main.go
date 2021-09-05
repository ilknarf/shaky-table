package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/ilknarf/shaky-table/api"
	"github.com/ilknarf/shaky-table/userdb"
	"github.com/rs/cors"
)

var corsAllowed []string

func init() {
	corsAllowed := strings.Split(os.Getenv("SHAKY_CORS_ALLOWED"), " ")
	if len(corsAllowed) == 0 {
		log.Fatal("missing ")
	}
}

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
	api.AddAPIRoutes(userDB, baseRouter)

	handler := getCorsHandler(baseRouter, corsAllowed)

	log.Println("Serving...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getCorsHandler(handler http.Handler, allowed []string) http.Handler {
	var c *cors.Cors

	if len(allowed) > 0 {
		c = cors.New(cors.Options{
			AllowedOrigins: allowed,
		})
	} else {
		c = cors.Default()
	}

	return c.Handler(handler)
}
