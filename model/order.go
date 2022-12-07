package model

import "time"

type Order struct {
	id           int
	typeOrder    string
	number       string
	customerName string `required:"true"`
	email        string `type:"email"`
	phone        string `type:"number"`
	date         time.Time
	quantity     int
	total        float64
	created_at   time.Time
	orderDetail  []OrderDetail
}

func (order *Order) GetId() *int {
	return &order.id
}

func (order *Order) GetType() *string {
	return &order.typeOrder
}
func (order *Order) GetNumber() *string {
	return &order.number
}

func (order *Order) GetCustomerName() *string {
	return &order.customerName
}

func (order *Order) GetEmail() *string {
	return &order.email
}
func (order *Order) GetPhone() *string {
	return &order.phone
}

func (order *Order) GetDate() *time.Time {
	return &order.date
}

func (order *Order) GetQuantity() *int {
	return &order.quantity
}

func (order *Order) GetTotal() *float64 {
	return &order.total
}

func (order *Order) GetCreatedAt() *time.Time {
	return &order.created_at
}

func (order *Order) GetOrderDetails() *[]OrderDetail {
	return &order.orderDetail
}

func (order *Order) SetId(id *int) {
	order.id = *id
}

func (order *Order) SetType(typeOrder *string) {
	order.typeOrder = *typeOrder
}
func (order *Order) SetNumber(number *string) {
	order.number = *number
}

func (order *Order) SetCustomerName(customerName *string) {
	order.customerName = *customerName
}

func (order *Order) SetEmail(email *string) {
	order.email = *email
}
func (order *Order) SetPhone(phone *string) {
	order.phone = *phone
}

func (order *Order) SetDate(date *time.Time) {
	order.date = *date
}

func (order *Order) SetQuantity(quantity *int) {
	order.quantity = *quantity
}

func (order *Order) SetTotal(total *float64) {
	order.total = *total
}

func (order *Order) SetCreatedAt(created_at *time.Time) {
	order.created_at = *created_at
}

func (order *Order) SetOrderDetails(orderDetail []OrderDetail) {
	order.orderDetail = orderDetail
}
