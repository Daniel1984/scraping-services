package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Streets []string

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetStreetNamesForUpdate(url string) Streets {
	req, _ := http.NewRequest(http.MethodGet, url+"/street-names-to-update", nil)
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	streetNames := Streets{}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &streetNames)

	return streetNames
}
