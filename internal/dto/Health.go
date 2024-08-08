package dto

import "time"

type Health struct {
	Environment string    `json:"environment"`
	Status      string    `json:"status"`
	Date        time.Time `json:"date"`
}
