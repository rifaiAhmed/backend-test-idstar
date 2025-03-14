package helpers

import "time"

func StringToTime(s string) (time.Time, error) {
	dateString := "2025-02-12"
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Now(), err
	}

	return parsedTime, nil
}

type Recipe struct {
	ID   uint    `json:"id"`
	Name string  `json:"name"`
	SKU  string  `json:"sku"`
	Cogs float64 `json:"cogs"`
}
