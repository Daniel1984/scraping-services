package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/scraping-service/listings-scraper/services"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	apiUrl := os.Getenv("API_URL")
	streets, err := services.GetStreetNamesForUpdate(apiUrl)

	if err != nil {
		log.Fatal("Err:", err)
	}

	fmt.Println("Streets to update -", len(streets))
	services.ScrapeListings(streets, apiUrl)
}
