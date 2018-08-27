package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"scraping-service/listings-scraper/utils"
	"time"
)

func GetPropertiesFromAirbnb(url string, ch chan []byte) {
	var httpClient = &http.Client{Timeout: 10 * time.Second}
	userAgentStr := utils.GetUserAgent()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	if res, err := httpClient.Do(req); err != nil {
		fmt.Println("getting properties errror: ", err)
		defer res.Body.Close()
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		ch <- body
	}
}
