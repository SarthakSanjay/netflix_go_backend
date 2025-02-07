package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/sarthaksanjay/netflix-go/db"
	"github.com/sarthaksanjay/netflix-go/router"
)

func main() {
	fmt.Println("Netflix Apis")
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	r := router.Router()
	db.ConnectDB()
	defer db.DisconnectDB()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // todo:in production use actual clien url
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow cookies and credentials
	})

	// wrapped router with cors middleware
	handler := c.Handler(r)

	fmt.Println("Server is getting started...")
	fmt.Println("Server listing on port 4000")

	log.Fatal(http.ListenAndServe(":4000", handler))
}
