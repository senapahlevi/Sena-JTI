package main

import (
	"fmt"
	"log"
	"net/http"

	"01-Login/platform/authenticator"
	"01-Login/platform/database"
	"01-Login/platform/router"

	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect success", db)
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}

	auth, err := authenticator.New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rtr := router.New(auth, db)

	log.Print("Server listening on http://localhost:3000/")
	if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}
}
