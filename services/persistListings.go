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
	if sob, err := json.Marshal(listings); err != nil {
		fmt.Println("ERROR!", err)
	} else {
		fmt.Println("OK!")

		var httpClient = &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest(http.MethodPost, apiUrl+"/persist-listings", bytes.NewBuffer(sob))
		req.Header.Set("Content-Type", "application/json")
		res, err := httpClient.Do(req)

		if err != nil {
			fmt.Println(err)
		}

		defer res.Body.Close()
	}
}
