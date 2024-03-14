package domain

import "time"

type Item struct {
	Site      string    `json:"site"`
	Id        string    `json:"id"`
	Price     float64   `json:"price"`
	StartTime time.Time `json:"start_time"`
	Category  string    `json:"category"`
	Currency  string    `json:"currency"`
	Seller    string    `json:"seller"`
}
