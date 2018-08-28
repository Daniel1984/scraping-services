package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"scraping-service/listing-availability-scraper/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		apiUrl := os.Getenv("API_URL")
		listingIds := services.GetListingIdsToUpdate(apiUrl)
		fmt.Println(listingIds)
	}
}
