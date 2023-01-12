package main

type ItemDto struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"amount"`
}

type OrderDto struct {
	ID    string    `json:"id"`
	Items []ItemDto `json:"items"`
	Total float64   `json:"total"`
}
