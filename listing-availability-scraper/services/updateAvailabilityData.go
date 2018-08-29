package services

import (
	"encoding/json"
	"fmt"
	"scraping-service/listing-availability-scraper/models"
	"scraping-service/listing-availability-scraper/utils"
)

func UpdateAvailabilityData(listingIds []int, apiUrl string) {
	channel := make(chan []byte)

	for _, listingId := range listingIds {
		fmt.Printf("Getting availabilities for listingId: %d\n", listingId)
		fmt.Println("-------------------------------------------------")

		availabilityUrl := utils.GetAvailabilityUrl(listingId)
		go ScrapeAvailabilities(availabilityUrl, channel)
		response := <-channel
		availabilities := models.Availabilities{}

		if err := json.Unmarshal(response, &availabilities); err != nil {
			fmt.Println("Error getting availabilities: ", err)
		} else {
			fmt.Println(availabilities.CalendarMonths[0].Days)
		}
	}
}
