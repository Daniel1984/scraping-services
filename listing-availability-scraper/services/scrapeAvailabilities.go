package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"scraping-service/listing-availability-scraper/models"
	"scraping-service/listing-availability-scraper/utils"
	"scraping-service/shared"
	"time"
)

func ScrapeAvailabilities(listingId int, ch chan models.AvailabilitiesToUpdate) {
	availabilityUrl := utils.GetAvailabilityUrl(listingId)
	fmt.Printf("Getting availabilities for listingId: %v\n", availabilityUrl)
	fmt.Println("-------------------------------------------------")

	var httpClient = &http.Client{Timeout: 10 * time.Second}
	userAgentStr := shared.GetUserAgent()
	req, _ := http.NewRequest(http.MethodGet, availabilityUrl, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	availabilitiesWithListingId := &models.AvailabilitiesToUpdate{ListingId: listingId}

	if res, err := httpClient.Do(req); err != nil {
		fmt.Println("getting availabilities errror: ", err)
		defer res.Body.Close()
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		availabilities := &models.Availabilities{}

		if err := json.Unmarshal(body, availabilities); err != nil {
			fmt.Println("Error getting availabilities: ", err)
		} else {
			days := []models.Day{}
			for _, val := range availabilities.CalendarMonths {
				days = append(days, val.Days...)
			}

			availabilitiesWithListingId.Availabilities = days
		}
	}

	ch <- *availabilitiesWithListingId
}
