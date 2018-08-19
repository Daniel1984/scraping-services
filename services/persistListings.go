package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"scraping-service/models"
	"time"
)

func PersistListings(listings models.Listings, apiUrl string) {
	jsonString, errMarshal := json.Marshal(listings)

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
}
