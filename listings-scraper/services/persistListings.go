package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listings-scraper/models"
	"net/http"
)

func PersistListings(listings models.Listings, apiUrl string) {
	if sob, err := json.Marshal(listings); err != nil {
		fmt.Println("Error: ", err)
	} else {
		req, err := http.Post(apiUrl+"/persist-listings", "application/json", bytes.NewBuffer(sob))

		if err != nil {
			fmt.Println(err)
		}

		defer req.Body.Close()
	}
}
