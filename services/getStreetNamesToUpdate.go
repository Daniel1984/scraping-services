package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"scraping-service/models"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetStreetNamesForUpdate(url string) models.Streets {
	req, _ := http.NewRequest(http.MethodGet, url+"/street-names-to-update", nil)
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer res.Body.Close()

	streetNames := models.Streets{}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &streetNames)

	return streetNames
}
