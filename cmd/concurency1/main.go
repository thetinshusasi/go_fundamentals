package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type OpenHours struct {
	Day   string `json:"day"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

type Hours struct {
	Day   string
	Open  string
	Close string
	ID    int
}

type Output struct {
	Days  string `json:"days"`
	Open  string `json:"open"`
	Close string `json:"close"`
}

func GroupOpenHours(input []OpenHours) []Output {
	weekEnum := map[string]int{"Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4, "Friday": 5, "Saturday": 6, "Sunday": 7}
	weekList := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	daysMaps := make(map[string]Hours)
	for _, oh := range input {
		daysMaps[oh.Day] = Hours{Day: oh.Day, Open: oh.Open, Close: oh.Close, ID: weekEnum[oh.Day]}
	}

	// Add missing days
	for _, day := range weekList {
		if _, exists := daysMaps[day]; !exists {
			daysMaps[day] = Hours{Day: day, Open: "", Close: "", ID: weekEnum[day]}
		}
	}

	// Create a slice from the map and sort it by ID
	updateInput := make([]Hours, 0, len(daysMaps))
	for _, h := range daysMaps {
		updateInput = append(updateInput, h)
	}
	sort.Slice(updateInput, func(i, j int) bool {
		return updateInput[i].ID < updateInput[j].ID
	})

	// Generate the output
	var output []Output
	firstPtr, secondPtr := 0, 1
	for secondPtr < len(updateInput)+1 {
		if secondPtr < len(updateInput) && updateInput[firstPtr].Open == updateInput[secondPtr].Open && updateInput[firstPtr].Close == updateInput[secondPtr].Close {
			secondPtr++
		} else {
			daysRange := updateInput[firstPtr].Day
			if firstPtr != secondPtr-1 {
				daysRange += "-" + updateInput[secondPtr-1].Day
			}
			output = append(output, Output{Days: daysRange, Open: updateInput[firstPtr].Open, Close: updateInput[firstPtr].Close})
			firstPtr = secondPtr
			secondPtr++
		}
	}

	return output
}

func main() {
	// Test case
	openHours := []map[string]string{
		{
			"day":   "Monday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Tuesday",
			"open":  "8:00 AM",
			"close": "6:00 PM",
		},
		{
			"day":   "Wednesday",
			"open":  "8:00 AM",
			"close": "5:00 PM",
		},
		{
			"day":   "Thursday",
			"open":  "8:00 AM",
			"close": "6:00 PM",
		},
	}
	var convertedOpenHours []OpenHours
	for _, oh := range openHours {
		convertedOpenHours = append(convertedOpenHours, OpenHours{
			Day:   oh["day"],
			Open:  oh["open"],
			Close: oh["close"],
		})
	}

	groupedOpenHours := GroupOpenHours(convertedOpenHours)

	jsonData, err := json.MarshalIndent(groupedOpenHours, "", "    ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON output
	fmt.Println(string(jsonData))

}
