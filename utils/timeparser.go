package utils

import (
	"fmt"
	"time"
)

func ParseTime(dateStr string) (*time.Time, error) {
	layout := "02/01/2006"
	parsedTime, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	start, _ := time.Parse(time.RFC3339, parsedTime.UTC().Format(time.RFC3339))
	return &start, err

}
