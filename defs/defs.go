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

// omitempty
