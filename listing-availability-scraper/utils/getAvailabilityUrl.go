package utils

import (
	"fmt"
	"time"
)

func GetAvailabilityUrl(listingId string) string {
	t := time.Now()

	return fmt.Sprintf("https://www.airbnb.com/api/v2/calendar_months"+
		"?key=d306zoyjsyarp7ifhu67rjxn52tv0t20"+
		"&currency=USD"+
		"&locale=en"+
		"&listing_id=%v"+
		"&month=%v"+
		"&year=%v"+
		"&count=4"+
		"&_format=with_conditions", listingId, int(t.Month())-1, int(t.Year()))
}
