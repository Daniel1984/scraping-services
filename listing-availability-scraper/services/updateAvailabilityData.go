package services

import (
	"fmt"
	"scraping-service/listing-availability-scraper/models"
)

func UpdateAvailabilityData(listingIds []int, apiUrl string) {
	availabilitiesChan := make(chan models.AvailabilitiesToUpdate)
	availabilitySubmitterChan := make(chan bool)

	for _, listingId := range listingIds {
		go ScrapeAvailabilities(listingId, availabilitiesChan)
	}

	for availabilities := range availabilitiesChan {
		go SubmitAvailabilities(availabilities, apiUrl, availabilitySubmitterChan)
	}

	for submitAvailabilitySuccess := range availabilitySubmitterChan {
		if submitAvailabilitySuccess == true {
			fmt.Println("SUBMIT OK")
		} else {
			fmt.Println("SUBMIT ERR")
		}
	}
}
