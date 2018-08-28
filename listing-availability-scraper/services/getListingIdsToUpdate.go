package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type listingIds []int

var httpClient = &http.Client{Timeout: 10 * time.Second}

func GetListingIdsToUpdate(url string) listingIds {
	req, _ := http.NewRequest(http.MethodGet, url+"/listings-to-update", nil)
	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer res.Body.Close()

	ids := listingIds{}
	body, _ := ioutil.ReadAll(res.Body)

	if err := json.Unmarshal(body, &ids); err != nil {
		return listingIds{}
	} else {
		return ids
	}
}
