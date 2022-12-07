package model

import "time"

type Product struct {
	id         int
	name       string
	price      float64
	stock      int
	created_at time.Time
}

func (product *Product) GetId() *int {
	return &product.id
}

func (product *Product) GetName() *string {
	return &product.name
}

func (product *Product) GetPrice() *float64 {
	return &product.price
}

func (product *Product) GetStock() *int {
	return &product.stock
}

func (product *Product) GetCreatedAt() *time.Time {
	return &product.created_at
}

func (product *Product) SetId(id *int) {
	product.id = *id
}

func (product *Product) SetName(name *string) {
	product.name = *name
}

func (product *Product) SetPrice(price *float64) {
	product.price = *price
}

func (product *Product) SetStock(stock *int) {
	product.stock = *stock
}

func (product *Product) SetCreateAt(created_at *time.Time) {
	product.created_at = *created_at
}
