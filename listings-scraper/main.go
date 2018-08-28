package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"scraping-service/listings-scraper/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiUrl := os.Getenv("API_URL")
	streets := services.GetStreetNamesForUpdate(apiUrl)
	fmt.Println(streets)
	services.GetListings(streets, apiUrl)
}
