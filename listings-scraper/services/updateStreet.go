package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listings-scraper/models"
	"net/http"
)

func UpdateStreet(street models.Street, apiUrl string) {
	if sob, err := json.Marshal(street); err != nil {
		fmt.Println("Error: ", err)
	} else {
		res, err := http.Post(apiUrl+"/update-street", "application/json", bytes.NewBuffer(sob))

		if err != nil {
			fmt.Println("Error: ", err)
		}

		defer res.Body.Close()
		fmt.Println("OK!")
	}
}
