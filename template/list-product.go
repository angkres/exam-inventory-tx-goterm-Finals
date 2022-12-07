package template

import (
	"database/sql"
	"exam-inventory/handler"
	"exam-inventory/helper"
	"fmt"

	"github.com/buger/goterm"
)

type orderTemplate struct {
	db           *sql.DB
	orderHandler handler.OrderHandler
}

func NewOrderTemplate(db *sql.DB, orderHandler handler.OrderHandler) *orderTemplate {
	return &orderTemplate{db, orderHandler}
}

func (template *orderTemplate) ListProduct() {
	helper.ClearScreen()

	products, err := template.orderHandler.GetProduct()
	if err != nil {
		panic(err)
	}

	box := goterm.NewBox(100|goterm.PCT, (len(products)+1)*2-(len(products)-1), 0)
	table := goterm.NewTable(0, 5, 1, ' ', 0)
	fmt.Fprintf(table, "ID\t| Name\t| Price\t| Stock\t| CreateAt\n")
	if len(products) == 0 {
		fmt.Fprintf(table, "Data kosong")
	} else {
		for _, v := range products {
			fmt.Fprintf(table, "%v\t| %s\t| Rp. %.f\t| %v\t| %v\n", *v.GetId(), *v.GetName(), *v.GetPrice(), *v.GetStock(), *v.GetCreatedAt())
		}
	}
	fmt.Fprint(box, table)
	fmt.Println(box)

	helper.BackHandler()
	Menu(template.db)
}
