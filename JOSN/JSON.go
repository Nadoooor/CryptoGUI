package JOSN

import (
	"encoding/json"
	"os"
)

// the History type that will has the structure of the data rows
type History struct {
	DateNtime string `json:"TIMESTAMP"`
	INPUT     string `json:"INPUT"`
	OUTPUT    string `json:"OUTPUT"`
	CIPHER    string `json:"CIPHER"`
}

func Save(things []History) {
	// just a function that takes the data and encode it to json and then save it to the file
	data, _ := json.MarshalIndent(things, "", "  ")
	os.WriteFile("history.json", data, 0644)
}

func Load() []History {
	// reading the data from the file and return the history array with all the data
	here, _ := os.ReadFile("history.json")
	var history []History
	json.Unmarshal(here, &history)
	return history

}
