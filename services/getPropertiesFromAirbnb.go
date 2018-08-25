package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func GetPropertiesFromAirbnb(url string, ch chan []byte) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	if res, err := httpClient.Do(req); err != nil {
		fmt.Println("Error gettign response body: ", err)
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		ch <- body
	}
}
