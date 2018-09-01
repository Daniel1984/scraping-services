package services

import (
	"fmt"
	"scraping-service/listing-availability-scraper/models"
	"scraping-service/listing-availability-scraper/utils"
)

func UpdateAvailabilityData(listingIds []int, apiUrl string) {
	channel := make(chan models.Availabilities)

	for _, listingId := range listingIds {
		availabilityUrl := utils.GetAvailabilityUrl(listingId)

		fmt.Printf("Getting availabilities for listingId: %v\n", availabilityUrl)
		fmt.Println("-------------------------------------------------")

		go ScrapeAvailabilities(availabilityUrl, channel)
		response := <-channel
		fmt.Println("got response ", response)
	}
}
