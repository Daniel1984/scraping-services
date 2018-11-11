package services

import (
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listings-scraper/models"
	"github.com/scraping-service/listings-scraper/utils"
	"math/rand"
	"net/url"
	"sort"
	"time"
)

func getRandIntInRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func ScrapeListings(streetNames []models.Street, apiUrl string) {
	channel := make(chan []byte)

	for _, street := range streetNames {
		if street.Name == "" {
			continue
		}

		streetName := &url.URL{Path: street.Name}
		encodedStreetName := streetName.String()
		listingUrl := utils.GetListingsURL(encodedStreetName)
		var sectionOffset int = 0
		var itemsOffset int = 0

		for {
			offsetUrl := fmt.Sprintf("%s%s%d%s%d", listingUrl, "&section_offset=", sectionOffset, "&items_offset=", itemsOffset)
			randReqDelay := getRandIntInRange(1, 2)
			time.Sleep(time.Duration(randReqDelay) * time.Second)
			go GetPropertiesFromAirbnb(offsetUrl, channel)
			responseBody := <-channel
			listing := models.Listing{}

			fmt.Printf("Street: %s, Offset: %d\n", street.Name, itemsOffset)

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
					itemsOffset = exploreTabs.PaginationMetadata.ItemsOffset
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
