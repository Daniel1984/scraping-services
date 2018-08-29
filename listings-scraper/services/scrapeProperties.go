package services

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"scraping-service/listings-scraper/models"
	"scraping-service/listings-scraper/utils"
	"sort"
	"time"
)

func getRandIntInRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func GetListings(streetNames models.Streets, apiUrl string) {
	channel := make(chan []byte)

	for _, street := range streetNames {
		streetName := &url.URL{Path: street.Name}
		encodedStreetName := streetName.String()
		listingUrl := utils.GetListingsURL(encodedStreetName)
		sectionOffset := 0

		for {
			offsetUrl := fmt.Sprintf("%s%s%d", listingUrl, "&section_offset=", sectionOffset)
			randReqDelay := getRandIntInRange(3, 5)
			time.Sleep(time.Duration(randReqDelay) * time.Second)
			go GetPropertiesFromAirbnb(offsetUrl, channel)
			responseBody := <-channel
			listing := models.Listing{}

			fmt.Printf("Street: %s, Offset: %d\n", street.Name, sectionOffset)
			fmt.Println("-------------------------------------------------")

			if err := json.Unmarshal(responseBody, &listing); err != nil {
				fmt.Println("Error unmarshal rsp body: ", err)
				UpdateStreet(street, apiUrl)
				continue
			} else {
				hometabIndex := sort.Search(len(listing.ExploreTabs), func(i int) bool {
					tabId := listing.ExploreTabs[i].TabId
					tabName := listing.ExploreTabs[i].TabName
					return tabId == "home_tab" || tabId == "all_tab" || tabName == "Homes"
				})

				if hometabIndex < len(listing.ExploreTabs) {
					exploreTabs := listing.ExploreTabs[hometabIndex]
					sectionOffset = exploreTabs.PaginationMetadata.SectionOffset
					PersistListings(exploreTabs.Sections[0].Listings, apiUrl)

					if exploreTabs.PaginationMetadata.HasNextPage == false {
						UpdateStreet(street, apiUrl)
						break
					}
				}
			}
		}
	}
}
