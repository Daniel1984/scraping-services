package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/scraping-service/listing-availability-scraper/services"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		apiUrl := os.Getenv("API_URL")

		listingIds, err := services.GetListingIdsToUpdate(apiUrl)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Listings to update -", len(listingIds))

		services.UpdateAvailabilityData(listingIds, apiUrl)
	}
}
