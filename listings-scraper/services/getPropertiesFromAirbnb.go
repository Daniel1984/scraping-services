package services

import (
	"fmt"
	"github.com/scraping-service/shared"
	"io/ioutil"
	"net/http"
)

func GetPropertiesFromAirbnb(url string, ch chan []byte) {
	httpClient := &http.Client{}
	userAgentStr := shared.GetUserAgent()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("getting properties errror: ", err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	ch <- body
}
