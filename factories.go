package main

import (
	"strconv"

	"github.com/google/uuid"
)

func makeItemFactory(quantity int) []ItemDto {
	var items []ItemDto
	for i := 0; i < quantity; i++ {
		items = append(items, ItemDto{
			ID:       uuid.New().String(),
			Name:     "Item " + strconv.Itoa(i),
			Price:    10.0,
			Quantity: 1,
		})
	}
	return items
}

func makeOrderFactory(quantity int) []OrderDto {
	var orders []OrderDto

	for i := 0; i < quantity; i++ {
		orders = append(orders, OrderDto{
			ID:    uuid.New().String(),
			Items: makeItemFactory(10),
			Total: 0,
		})
	}

	return orders
}
