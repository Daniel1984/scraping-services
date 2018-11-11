package services

import (
	"encoding/json"
	"fmt"
	"github.com/scraping-service/listings-scraper/models"
	"io/ioutil"
	"net/http"
)

func GetStreetNamesForUpdate(url string) ([]models.Street, error) {
	req, err := http.Get(url + "/street-names-to-update")

	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	defer req.Body.Close()

	streetNames := []models.Street{}
	body, _ := ioutil.ReadAll(req.Body)

	if err := json.Unmarshal(body, &streetNames); err != nil {
		return nil, err
	} else {
		return streetNames, nil
	}
}
