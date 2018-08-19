package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Street struct {
	Name string `json:"name"`
	Id   string `json:"_id"`
}

func UpdateStreet(street Street, apiUrl string) {
	jsonString, errMarshal := json.Marshal(street)

	if errMarshal != nil {
		fmt.Println("ERROR!")
	} else {
		fmt.Println("OK!")
		var httpClient = &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest(http.MethodPost, apiUrl+"/update-street", bytes.NewBuffer(jsonString))
		req.Header.Set("Content-Type", "application/json")
		res, err := httpClient.Do(req)

		if err != nil {
			fmt.Println("Error")
		}

		defer res.Body.Close()
	}
}
