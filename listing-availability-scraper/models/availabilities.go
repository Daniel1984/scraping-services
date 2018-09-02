package models

type AvailabilitiesToUpdate struct {
	ListingId      int
	Availabilities []Day
}

type Day struct {
	Available           bool   `json:"available"`
	Date                string `json:"date"`
	AvailableForCheckin bool   `json:"available_for_checkin"`
	Price               struct {
		Date                string  `json:"date"`
		LocalAdjustedPrice  int     `json:"local_adjusted_price"`
		LocalCurrency       string  `json:"local_currency"`
		LocalPrice          int     `json:"local_price"`
		NativeAdjustedPrice int     `json:"native_adjusted_price"`
		NativeCurrency      string  `json:"native_currency"`
		NativePrice         float64 `json:"native_price"`
		Type                string  `json:"type"`
		LocalPriceFormatted string  `json:"local_price_formatted"`
	} `json:"price"`
}

type Availabilities struct {
	CalendarMonths []struct {
		Days  []Day `json:"days"`
		Month int   `json:"month"`
		Year  int   `json:"year"`
	} `json:"calendar_months"`
}
