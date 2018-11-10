package services

import (
	"fmt"
	"github.com/scraping-service/listing-availability-scraper/models"
)

func UpdateAvailabilityData(listingIds []int, apiUrl string) {
	availabilitiesConsumer := getListingavailabilityProducer(listingIds)

	for availabilities := range availabilitiesConsumer {
		go SubmitAvailabilities(availabilities, apiUrl)
	}
}

func getListingavailabilityProducer(listingIds []int) <-chan models.AvailabilitiesToUpdate {
	ch := make(chan models.AvailabilitiesToUpdate)

	go func() {
		defer close(ch)
		for _, listingId := range listingIds {
			fmt.Println(listingId)
			ScrapeAvailabilities(listingId, ch)
		}
	}()

	return ch
}
