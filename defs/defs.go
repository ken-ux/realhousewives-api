package defs

import (
	"time"
)

type Series struct {
	Series_id     string     `json:"series_id"`
	Title         string     `json:"title"`
	Location      string     `json:"location"`
	Premiere_date *time.Time `json:"premiere_date"`
}

type Season struct {
	Season_id     int        `json:"-"`
	Series_id     string     `json:"series_id"`
	Season_number int        `json:"season_number"`
	Premiere_date *time.Time `json:"premiere_date"`
	Finale_date   *time.Time `json:"finale_date"`
}

// omitempty
