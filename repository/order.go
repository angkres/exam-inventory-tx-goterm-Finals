package repository

import (
	"context"
	"database/sql"
	"errors"
	"exam-inventory/helper"
	"exam-inventory/model"
)

type OrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, order model.Order) (model.Order, error)
	GetOrderByOrderNumber(ctx context.Context, typeOrder string, number string) (model.Order, error)
	FindAll(ctx context.Context) ([]model.Order, error)
}

type orderRepository struct {
	db *sql.DB
 }

func NewOrderRepository(db *sql.DB) *orderRepository {
	return &orderRepository{db}
}

func (repo *orderRepository) FindAll(ctx context.Context) ([]model.Order, error) {
	var query string = "SELECT id, type, number, customer_name, email, phone, date, quantity, total, created_at FROM orders"
	var orders []model.Order

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return orders, err
	}
	for rows.Next() {
		var order model.Order
		var email sql.NullString
		var phone sql.NullString
		rows.Scan(order.GetId(), order.GetType(), order.GetNumber(), order.GetCustomerName(), &email, &phone, order.GetDate(), order.GetQuantity(), order.GetTotal(), order.GetCreatedAt())
		emailString := helper.NullStringtoStringEmail(&email)
		order.SetEmail(&emailString)

		phoneString := helper.NullStringtoStringPhone(&phone)
		order.SetPhone(&phoneString)

		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *orderRepository) GetOrderByOrderNumber(ctx context.Context, typeOrder string, number string) (model.Order, error) {
	var query string = "SELECT id, type, number, customer_name, email, phone, date, quantity, total, created_at FROM orders WHERE type=? AND number=?"
	var order model.Order
	var email sql.NullString
	var phone sql.NullString

	rows := repo.db.QueryRowContext(ctx, query, typeOrder, number)
	err := rows.Scan(order.GetId(), order.GetType(), order.GetNumber(), order.GetCustomerName(), &email, &phone, order.GetDate(), order.GetQuantity(), order.GetTotal(), order.GetCreatedAt())

	emailString := helper.NullStringtoStringEmail(&email)
	order.SetEmail(&emailString)

	phoneString := helper.NullStringtoStringPhone(&phone)
	order.SetPhone(&phoneString)

	if err != nil {
		return order, errors.New("order number tidak ada dalam tabel orders")
	}
	return order, nil
}

func (repo *orderRepository) Insert(ctx context.Context, tx *sql.Tx, order model.Order) (model.Order, error) {
	//var query string = "INSERT INTO orders(type, number, customer_name, email, phone, date, quantity, total) VALUES(?,?,?,?,?,?,?,?)"
	var query string = "INSERT INTO orders(type, number, customer_name, email, phone, date, quantity, total) VALUES(?,?,?,NULLIF(?,''),NULLIF(?,''),?,?,?)"

	res, err := tx.ExecContext(ctx, query, order.GetType(), order.GetNumber(), order.GetCustomerName(), order.GetEmail(), order.GetPhone(), order.GetDate(), order.GetQuantity(), order.GetTotal())
	if err != nil {
		return order, err
	}
	lastInsertId, _ := res.LastInsertId()
	id := int(lastInsertId)
	order.SetId(&id)

	return order, nil
}
