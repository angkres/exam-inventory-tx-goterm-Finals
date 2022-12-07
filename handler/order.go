package handler

import (
	"context"
	"database/sql"
	"errors"
	"exam-inventory/helper"
	"exam-inventory/model"
	"exam-inventory/repository"
	"time"
)

type OrderHandler interface {
	GetProduct() ([]model.Product, error)
	InsertOrder(customerName, email, phone, typeOrder string, order []map[string]interface{}) (model.Order, error)
	GetOrderByOrderNumber(orderNumber string) (model.Order, error)
	ListOrder() ([]model.Order, error)
}

type orderHandler struct {
	db                    *sql.DB
	orderRepository       repository.OrderRepository
	orderDetailRepository repository.OrderDetailRepository
	productRepository     repository.ProductRepository
}

func NewOrderHandler(db *sql.DB, orderRepository repository.OrderRepository, orderDetailRepository repository.OrderDetailRepository, productRepository repository.ProductRepository) *orderHandler {
	return &orderHandler{db, orderRepository, orderDetailRepository, productRepository}
}

func (handler *orderHandler) ListOrder() ([]model.Order, error) {
	ctx := context.Background()
	orders, err := handler.orderRepository.FindAll(ctx)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (handler *orderHandler) GetOrderByOrderNumber(orderNumber string) (model.Order, error) {
	ctx := context.Background()
	var order model.Order

	typeOrder, number := helper.SplitString(orderNumber)

	order, err := handler.orderRepository.GetOrderByOrderNumber(ctx, typeOrder, number)
	if err != nil {
		return order, err
	}

	orderDetails, err2 := handler.orderDetailRepository.GetOrderDetailsByOrderId(ctx, *order.GetId())
	if err2 != nil {
		return order, err2
	}

	order.SetOrderDetails(orderDetails)
	return order, nil
}

func (handler *orderHandler) GetProduct() ([]model.Product, error) {
	ctx := context.Background()

	products, err := handler.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (handler *orderHandler) InsertOrder(customerName, email, phone, typeOrder string, orderDetails []map[string]interface{}) (model.Order, error) {
	ctx := context.Background()
	var order model.Order
	var orderDetailDatas []model.OrderDetail

	number := helper.RandomNumber()
	date := time.Now()

	order.SetType(&typeOrder)
	order.SetNumber(&number)
	order.SetDate(&date)

	order.SetCustomerName(&customerName)
	order.SetEmail(&email)
	order.SetPhone(&phone)

	var tempTotal float64
	var tempQuantity int

	tx, err0 := handler.db.BeginTx(ctx, nil)
	defer tx.Commit()
	if err0 != nil {
		return order, err0
	}
	for _, v := range orderDetails {
		var orderDetail model.OrderDetail
		productName := v["productName"].(string)
		quantity := v["quantity"].(int)

		orderDetail.SetProductName(&productName)
		orderDetail.SetQuantity(&quantity)

		product, err := handler.productRepository.GetProductByProductName(ctx, productName)
		if err != nil {
			return order, err
		}
		orderDetail.SetProductId(product.GetId())
		orderDetail.SetPrice(product.GetPrice())

		var price float64 = *product.GetPrice()
		var quantityFloat float64 = float64(quantity)
		var total = price * quantityFloat
		orderDetail.SetTotal(&total)
		tempTotal += total
		tempQuantity += quantity

		orderDetailDatas = append(orderDetailDatas, orderDetail)

		stock := *product.GetStock()
		if *order.GetType() == "PO" {
			updateStock := stock + quantity
			product.SetStock(&updateStock)
		} else {
			updateStock := stock - quantity
			if updateStock < 0 {
				return order, errors.New("stock tidak cukup untuk dijual")
			}
			product.SetStock(&updateStock)
		}

		err2 := handler.productRepository.Update(ctx, tx, product)
		if err2 != nil {
			tx.Rollback()
			return order, err2
		}
	}
	order.SetQuantity(&tempQuantity)
	order.SetTotal(&tempTotal)

	order, err := handler.orderRepository.Insert(ctx, tx, order)
	if err != nil {
		tx.Rollback()
		return order, err
	}

	orderDetailsEnd, err2 := handler.orderDetailRepository.InsertOrderDetails(ctx, tx, orderDetailDatas, *order.GetId())
	if err2 != nil {
		tx.Rollback()
		return order, err2
	}

	order.SetOrderDetails(orderDetailsEnd)
	return order, nil
}
