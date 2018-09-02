package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"scraping-service/listing-availability-scraper/models"
	"time"
)

func SubmitAvailabilities(availabilities models.AvailabilitiesToUpdate, apiUrl string, c chan bool) {
	fmt.Println(availabilities)
	if sob, err := json.Marshal(availabilities); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("OK!")
		var httpClient = &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest(http.MethodPost, apiUrl+"/update-availabilities", bytes.NewBuffer(sob))
		req.Header.Set("Content-Type", "application/json")

		if res, err := httpClient.Do(req); err != nil {
			fmt.Println("Error: ", err)
			c <- true
			defer res.Body.Close()
		} else {
			c <- false
			defer res.Body.Close()
		}
	}
}
