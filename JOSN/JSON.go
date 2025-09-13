package JOSN

import (
	"encoding/json"
	"os"
)

type History struct {
	DateNtime string `json:"TIMESTAMP"`
	INPUT     string `json:"INPUT"`
	OUTPUT    string `json:"OUTPUT"`
	CIPHER    string `json:"CIPHER"`
}

func Save(things []History) {

	data, _ := json.MarshalIndent(things, "", "  ")
	os.WriteFile("history.json", data, 0644)
	println("saved")
}

func Load() []History {
	here, _ := os.ReadFile("history.json")
	var history []History
	json.Unmarshal(here, &history)
	println("loaded")
	return history

}
