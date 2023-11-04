package entities

import "time"

type Feed struct {
	ID        uint      `json:"id"`
	Body      string    `json:"body"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"createdAt"`
}
