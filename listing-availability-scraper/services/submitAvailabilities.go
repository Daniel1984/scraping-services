package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listing-availability-scraper/models"
	"net/http"
)

func SubmitAvailabilities(availabilities models.AvailabilitiesToUpdate, apiUrl string) {
	if sob, err := json.Marshal(availabilities); err != nil {
		fmt.Println("Error: ", err)
	} else {
		jsonByte := []byte(sob)

		res, err := http.Post(apiUrl+"/update-availabilities", "application/json", bytes.NewBuffer(jsonByte))
		if err != nil {
			fmt.Println("Err:", err)
		}

		defer res.Body.Close()
	}
}
