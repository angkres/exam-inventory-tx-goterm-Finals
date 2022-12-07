package repository

import (
	"context"
	"database/sql"
	"exam-inventory/model"
)

type OrderDetailRepository interface {
	InsertOrderDetails(ctx context.Context, tx *sql.Tx, orderDetails []model.OrderDetail, orderId int) ([]model.OrderDetail, error)
	GetOrderDetailsByOrderId(ctx context.Context, orderId int) ([]model.OrderDetail, error)
}

type orderDetailRepository struct {
	db *sql.DB
}

func NewOrderDetailRepository(db *sql.DB) *orderDetailRepository {
	return &orderDetailRepository{db}
}

func (repo *orderDetailRepository) InsertOrderDetails(ctx context.Context, tx *sql.Tx, orderDetails []model.OrderDetail, orderId int) ([]model.OrderDetail, error) {
	var query string = "INSERT INTO order_details(order_id, product_id, product_name, price, quantity, total) VALUES(?,?,?,?,?,?)"

	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, v := range orderDetails {
		res, err := stmt.ExecContext(ctx, orderId, v.GetProductId(), v.GetProductName(), v.GetPrice(), v.GetQuantity(), v.GetTotal())
		if err != nil {
			return nil, err
		}
		lastInsertId, _ := res.LastInsertId()
		id := int(lastInsertId)
		v.SetId(&id)
		v.SetOrderId(&orderId)
	}
	return orderDetails, nil
}

func (repo *orderDetailRepository)GetOrderDetailsByOrderId(ctx context.Context, orderId int) ([]model.OrderDetail, error){
	var orderDetailDatas []model.OrderDetail
	var query string = "SELECT id, order_id, product_id, product_name, price, quantity, total, created_at FROM order_details WHERE order_id=?"

	rows, err := repo.db.QueryContext(ctx, query, orderId)
	if err != nil {
		return orderDetailDatas, err
	}
	for rows.Next() {
		var orderDetailData model.OrderDetail
		rows.Scan(orderDetailData.GetId(), orderDetailData.GetOrderId(), orderDetailData.GetProductId(), orderDetailData.GetProductName(), orderDetailData.GetPrice(), orderDetailData.GetQuantity(), orderDetailData.GetTotal(), orderDetailData.GetCreatedAt())
		orderDetailDatas = append(orderDetailDatas, orderDetailData)
	}
	return orderDetailDatas, err
}
