package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// Author : Describes composer
	type Author struct {
		Name     string `json:"composer_name"`
		LastName string `json:"composer_last_name"`
	}

	// Song : describes a song
	type Song struct {
		*Author
		Band          string    `json:"band"`
		Song          string    `json:"song"`
		ReleaseDate   time.Time `json:"release_date"`
		Label         string    `json:"label,omitempty"` //This is weird they go inside the ""
		ChartPosition int       `json:"chart_position"`
		Producer      string    `json:"-"` //allways omitted
	}

	song := &Song{
		Author: &Author{Name: "Bob",
			LastName: "Dylan"},
		Band:        "The Rolling Stones",
		Song:        "Like a rolling stone",
		ReleaseDate: time.Date(1995, 7, 13, 0, 0, 0, 0, time.UTC),
	}

	jsonSong, _ := json.Marshal(song)
	fmt.Println(string(jsonSong))
}
