package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/hktrib/simple_bank/cmd/api"
)

func main() {
	// Loading env variables for
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	srv := api.NewServer(os.Getenv("SENDER_EMAIL"), os.Getenv("SENDER_PASSCODE"))
	srv.MountHandlers()
	go func() {

		// Starting the server
		http.ListenAndServe(":8080", srv.Router)
	}()
	select {}
}
