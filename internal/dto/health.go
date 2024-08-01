package dto

import "time"

type Health struct {
	Status string    `json:"status"`
	Date   time.Time `json:"date"`
}
