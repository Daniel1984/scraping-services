package services

import (
	"encoding/json"
	"fmt"
	"net/url"
	"scraping-service/models"
	"sort"
	"time"
)

func GetListings(streetNames models.Streets, apiUrl string) {
	channel := make(chan []byte)

	for _, street := range streetNames {
		streetName := &url.URL{Path: street.Name}
		encodedStreetName := streetName.String()
		listingUrl := getListingsURL(encodedStreetName)
		sectionOffset := 0

		for {
			offsetUrl := fmt.Sprintf("%s%s%d", listingUrl, "&section_offset=", sectionOffset)
			time.Sleep(1 * time.Second)
			go GetPropertiesFromAirbnb(offsetUrl, channel)
			responseBody := <-channel
			listing := models.Listing{}

			fmt.Println("-------------------------------------------------")
			fmt.Printf("Street: %s, Offset: %d\n", street.Name, sectionOffset)
			fmt.Println("-------------------------------------------------")

			if err := json.Unmarshal(responseBody, &listing); err != nil {
				fmt.Println("Error unmarshal rsp body: ", err)
				UpdateStreet(street, apiUrl)
				break
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
