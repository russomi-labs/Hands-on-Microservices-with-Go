package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Song : type to describe a Song
type Song struct {
	Band          string    `json:"band"`
	Song          string    `json:"song"`
	ReleaseDate   time.Time `json:"release_date"`
	Label         string    `json:"label,omitempty"` //omitempty goes inside the ""
	ChartPosition int       `json:"chart_position"`
	Producer      string    `json:"-"` //always omitted
}

func main() {
	song := &Song{
		Band:        "Blur",
		Song:        "Coffee & Tv",
		ReleaseDate: time.Date(1999, 6, 28, 0, 0, 0, 0, time.UTC),
		Producer:    "William Orbit",
	}

	jsonSong, _ := json.Marshal(song)
	fmt.Println(string(jsonSong))
}
