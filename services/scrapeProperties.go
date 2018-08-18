package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"scraping-service/models"
	"sort"
	"time"
)

func getListingsURL(location string) string {
	return fmt.Sprintf("https://www.airbnb.com/api/v2/explore_tabs"+
		"?version=1.3.9"+
		"client_session_id=4bb22f21-bc3b-4399-89f1-dd1fb7d3482b"+
		"&_format=for_explore_search_web"+
		"&experiences_per_grid=20"+
		"&items_per_grid=18"+
		"&guidebooks_per_grid=0"+
		"&auto_ib=false"+
		"&fetch_filters=true"+
		"&is_guided_search=false"+
		"&is_new_trips_cards_experiment=true"+
		"&is_new_homes_cards_experiment=false"+
		"&luxury_pre_launch=false"+
		"&screen_size=small"+
		"&show_groupings=false"+
		"&supports_for_you_v3=true"+
		"&timezone_offset=180"+
		"&metadata_only=false"+
		"&is_standard_search=true"+
		"&selected_tab_id=home_tab"+
		"&tab_id=home_tab"+
		"&location=%v"+
		"&_intents=p1"+
		"&key=d306zoyjsyarp7ifhu67rjxn52tv0t20"+
		"&currency=USD"+
		"&locale=en", location)
}

func getResponseBody(url string, ch chan []byte) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("Error gettign response body: ", err)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	ch <- body
}

// Listing root path handler
func GetListings(streetNames []string, apiUrl string) {
	var listingUrls = []string{}

	for _, street := range streetNames {
		streetName := &url.URL{Path: street}
		encodedStreetName := streetName.String()
		url := getListingsURL(encodedStreetName)
		listingUrls = append(listingUrls, url)
	}

	channel := make(chan []byte)

	for _, listingUrl := range listingUrls {
		sectionOffset := 1

		for {
			offsetUrl := fmt.Sprintf("%s%s%d", listingUrl, "&section_offset=", sectionOffset)
			go getResponseBody(offsetUrl, channel)
			responseBody := <-channel

			listing := models.Listing{}
			json.Unmarshal(responseBody, &listing)

			hometabIndex := sort.Search(len(listing.ExploreTabs), func(i int) bool {
				tabId := listing.ExploreTabs[i].TabId
				tabName := listing.ExploreTabs[i].TabName
				return tabId == "home_tab" || tabId == "all_tab" || tabName == "Homes"
			})

			exploreTabs := listing.ExploreTabs[hometabIndex]
			sectionOffset = exploreTabs.PaginationMetadata.SectionOffset

			jsonString, errMarshal := json.Marshal(exploreTabs.Sections[0].Listings)
			if errMarshal != nil {
				fmt.Println("ERROR!")
			} else {
				fmt.Println("OK!")

				var httpClient = &http.Client{Timeout: 10 * time.Second}
				req, _ := http.NewRequest(http.MethodPost, apiUrl+"/persist-listings", bytes.NewBuffer(jsonString))
				req.Header.Set("Content-Type", "application/json")
				res, err := httpClient.Do(req)

				if err != nil {
					fmt.Println(err)
				}

				defer res.Body.Close()
			}

			if exploreTabs.PaginationMetadata.HasNextPage == false {
				break
			}
		}
	}
}
