package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type listingIds []int

func GetListingIdsToUpdate(url string) (listingIds, error) {
	res, err := http.Get(url + "/listings-to-update")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	ids := listingIds{}

	if err := json.Unmarshal(body, &ids); err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	} else {
		return ids, nil
	}
}
