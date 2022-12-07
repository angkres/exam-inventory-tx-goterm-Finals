package model

import "time"

type OrderDetail struct {
	id          int
	orderId     int
	productId   int
	productName string
	price       float64
	quantity    int
	total       float64
	created_at  time.Time
}

func (orderDetail *OrderDetail) GetId() *int {
	return &orderDetail.id
}

func (orderDetail *OrderDetail) GetOrderId() *int {
	return &orderDetail.orderId
}

func (orderDetail *OrderDetail) GetProductId() *int {
	return &orderDetail.productId
}

func (orderDetail *OrderDetail) GetProductName() *string {
	return &orderDetail.productName
}

func (orderDetail *OrderDetail) GetPrice() *float64 {
	return &orderDetail.price
}

func (orderDetail *OrderDetail) GetQuantity() *int {
	return &orderDetail.quantity
}

func (orderDetail *OrderDetail) GetTotal() *float64 {
	return &orderDetail.total
}

func (orderDetail *OrderDetail) GetCreatedAt() *time.Time {
	return &orderDetail.created_at
}

func (orderDetail *OrderDetail) SetId(id *int) {
	orderDetail.id = *id
}

func (orderDetail *OrderDetail) SetOrderId(orderId *int) {
	orderDetail.orderId = *orderId
}

func (orderDetail *OrderDetail) SetProductId(productId *int) {
	orderDetail.productId = *productId
}

func (orderDetail *OrderDetail) SetProductName(productName *string) {
	orderDetail.productName = *productName
}

func (orderDetail *OrderDetail) SetPrice(price *float64) {
	orderDetail.price = *price
}

func (orderDetail *OrderDetail) SetQuantity(quantity *int) {
	orderDetail.quantity = *quantity
}

func (orderDetail *OrderDetail) SetTotal(total *float64) {
	orderDetail.total = *total
}

func (orderDetail *OrderDetail) SetCreateAt(created_at *time.Time) {
	orderDetail.created_at = *created_at
}
