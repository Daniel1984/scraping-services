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
	if sob, err := json.Marshal(street); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("OK!")
		var httpClient = &http.Client{Timeout: 10 * time.Second}
		req, _ := http.NewRequest(http.MethodPost, apiUrl+"/update-street", bytes.NewBuffer(sob))
		req.Header.Set("Content-Type", "application/json")
		res, err := httpClient.Do(req)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		defer res.Body.Close()
	}
}
