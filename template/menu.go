package template

import (
	"database/sql"
	"exam-inventory/handler"
	"exam-inventory/helper"
	"exam-inventory/repository"
	"fmt"
	"os"
)

func Menu(db *sql.DB) {
	// Dependency Injection
	orderRepository := repository.NewOrderRepository(db)
	orderDetailRepository := repository.NewOrderDetailRepository(db)
	productRepository := repository.NewProductRepository(db)
	orderHandler := handler.NewOrderHandler(db, orderRepository, orderDetailRepository, productRepository)
	orderTemplate := NewOrderTemplate(db, orderHandler)

	helper.ClearScreen()
	fmt.Println("Menu")
	fmt.Println("=================")
	fmt.Println("1. List Product")
	fmt.Println("2. Insert Purchases Order")
	fmt.Println("3. Insert Sales Order")
	fmt.Println("4. List Order By Order Number")
	fmt.Println("5. Exit")

	var menu int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		orderTemplate.ListProduct()

	case 2:
		orderTemplate.AddPOrder()
	case 3:
		orderTemplate.AddSOrder()
	case 4:
		orderTemplate.ListOrderByNumberOrder()
	case 5:
		os.Exit(0)
	default:
		Menu(db)
	}
}
