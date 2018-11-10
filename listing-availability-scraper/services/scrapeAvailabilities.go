package services

import (
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listing-availability-scraper/models"
	"github.com/scraping-service/listing-availability-scraper/utils"
	"github.com/scraping-service/shared"
	"io/ioutil"
	"net/http"
)

func ScrapeAvailabilities(listingId int, ch chan<- models.AvailabilitiesToUpdate) {
	availabilityUrl := utils.GetAvailabilityUrl(listingId)

	var httpClient = &http.Client{}
	userAgentStr := shared.GetUserAgent()
	req, _ := http.NewRequest(http.MethodGet, availabilityUrl, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	availabilitiesWithListingId := models.AvailabilitiesToUpdate{ListingId: listingId}
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("getting availabilities errror: ", err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	availabilities := &models.Availabilities{}

	if err := json.Unmarshal(body, availabilities); err != nil {
		fmt.Println("Error getting availabilities: ", err)
	} else {
		for _, val := range availabilities.CalendarMonths {
			availabilitiesWithListingId.Availabilities = append(availabilitiesWithListingId.Availabilities, val.Days...)
		}

		ch <- availabilitiesWithListingId
	}
}
