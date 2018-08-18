package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"scraping-service/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiUrl := os.Getenv("API_URL")
	streets := services.GetStreetNamesForUpdate(apiUrl)
	services.GetListings(streets, apiUrl)
}
