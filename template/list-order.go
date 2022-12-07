package template

import (
	"exam-inventory/helper"
	"fmt"

	"github.com/buger/goterm"
)

func (template *orderTemplate) ListOrderByNumberOrder() {
	helper.ClearScreen()
	orders, err := template.orderHandler.ListOrder()
	if err != nil {
		panic(err)
	}
	box := goterm.NewBox(100|goterm.PCT, (len(orders)+1)*2-(len(orders)-1), 0)
	table := goterm.NewTable(0, 5, 1, ' ', 0)
	fmt.Fprintf(table, "Daftar Order Number\n")
	for i, v := range orders {
		fmt.Fprintf(table, "%v. %s-%s\n", i+1, *v.GetType(), *v.GetNumber())
	}
	fmt.Fprint(box, table)
	fmt.Println(box)
	var orderNumber string
	fmt.Print("Input Order Number yang akan dicari: ")
	fmt.Scanln(&orderNumber)
	order, err := template.orderHandler.GetOrderByOrderNumber(orderNumber)
	if err != nil {
		panic(err)
	}
	template.Receipt(order)
}
