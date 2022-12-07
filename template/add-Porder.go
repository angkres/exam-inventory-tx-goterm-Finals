package template

import (
	"bufio"
	"exam-inventory/helper"
	"exam-inventory/model"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/buger/goterm"
)

func (template *orderTemplate) AddPOrder() {
	helper.ClearScreen()
	var nameCustomer, email, phone, withEmail, withPhone string
	var typeOrder string = "PO"

	fmt.Println("Add Purchases Order - Customer")
	fmt.Println("====================")
	template.InputName(&nameCustomer)
	fmt.Print("Apakah anda akan menginput email (y/n): ")
	fmt.Scanln(&withEmail)
	if strings.ToLower(withEmail) == "y" {
		template.InputEmail(&email)
	}
	fmt.Print("Apakah anda akan menginput phone (y/n): ")
	fmt.Scanln(&withPhone)
	if strings.ToLower(withPhone) == "y" {
		template.InputPhone(&phone)
	}

	helper.PauseHandler()
	template.InputOrders(&nameCustomer, &email, &phone, &typeOrder)
}

func (template *orderTemplate) Receipt(order model.Order) {
	helper.ClearScreen()

	table := goterm.NewTable(0, 5, 1, ' ', 0)
	fmt.Fprintf(table, "Nomer Transaksi: %s-%s\n", *order.GetType(), *order.GetNumber())
	fmt.Fprintf(table, "Nama Customer: %s\n", *order.GetCustomerName())
	fmt.Fprintf(table, "Email: %s\n", *order.GetEmail())
	fmt.Fprintf(table, "Phone: %s\n", *order.GetPhone())
	fmt.Fprintf(table, "Date: %s\n", helper.DateToString(*order.GetDate()))

	order_details := *order.GetOrderDetails()
	box := goterm.NewBox(100|goterm.PCT, 13+len(order_details)*2, 0)
	fmt.Fprintf(table, "\n")
	fmt.Fprintf(table, "No.\t| Name\t| Price\t| Quantity\t| Total\n")
	for i, v := range order_details {
		fmt.Fprintf(table, "%v\t| %s\t| Rp. %.f\t| %v\t| Rp. %.f\n", i+1, *v.GetProductName(), *v.GetPrice(), *v.GetQuantity(), *v.GetTotal())
	}

	fmt.Fprintf(table, "\n")
	fmt.Fprintf(table, "Grand Total\t: Rp. %.f\n", *order.GetTotal())
	fmt.Fprintf(table, "Grand Quantity\t: %v\n", *order.GetQuantity())
	fmt.Fprint(box, table)
	fmt.Println(box)
	helper.BackHandler()
	Menu(template.db)
}

func (template *orderTemplate) InputOrders(nameCustomer, email, phone, typeOrder *string) {
	var orders []map[string]interface{}
	template.InputOrder(&orders)
	order, err := template.orderHandler.InsertOrder(*nameCustomer, *email, *phone, *typeOrder, orders)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Anda perlu input produk dari awal lagi.")
		helper.PauseHandler()
		template.InputOrders(nameCustomer, email, phone, typeOrder)
	}
	template.Receipt(order)
}

func (template *orderTemplate) InputOrder(orders *[]map[string]interface{}) {
	helper.ClearScreen()
	fmt.Println("Form Order")
	fmt.Println("==================")
	var lanjut string
	var quantity int

	fmt.Print("Nama Produk: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	productName := scanner.Text()

	template.InputQuantity(&quantity)

	order := make(map[string]interface{})
	order = map[string]interface{}{
		"productName": productName,
		"quantity":    quantity,
	}
	*orders = append(*orders, order)

	fmt.Print("Input product lagi: ")
	fmt.Scanln(&lanjut)
	if strings.ToLower(lanjut) == "y" {
		template.InputOrder(orders)
	}
}

func (template *orderTemplate) InputName(nameCustomer *string) {
	fmt.Print("Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if !ValidateName(&input) {
		fmt.Println("Name tidak boleh kosong")
		template.InputName(&input)
	}
	*nameCustomer = input
}

func (template *orderTemplate) InputQuantity(quantity *int) {
	var input int
	fmt.Print("Quantity: ")
	fmt.Scanln(&input)

	if input <= 0 {
		fmt.Println("Quantity tidak boleh kosong atau negatif")
		template.InputQuantity(&input)
	}
	*quantity = input
}

func (template *orderTemplate) InputEmail(email *string) {
	var input string
	fmt.Print("Email: ")
	fmt.Scanln(&input)

	if !ValidateEmail(&input) {
		fmt.Println("Email harus mengandung @ tidak boleh kosong")
		template.InputEmail(&input)
	}
	*email = input
}

func (template *orderTemplate) InputPhone(phone *string) {
	var input string
	fmt.Print("Phone: ")
	fmt.Scanln(&input)

	if !ValidatePhone(&input) {
		fmt.Println("Phone harus angka tidak boleh kosong")
		template.InputPhone(&input)
	}
	*phone = input
}

func ValidateName(name *string) bool {
	var c model.Order
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(3).Tag.Get("required") == "true" {
		if *name == "" {
			return false
		}
	}
	return true
}

func ValidateEmail(email *string) bool {
	var c model.Order
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(4).Tag.Get("type") == "email" {
		if !strings.Contains(*email, "@") {
			return false
		}
	}
	return true
}

func ValidatePhone(phone *string) bool {
	var c model.Order
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(5).Tag.Get("type") == "number" {
		for _, v := range *phone {
			if v < 48 || v > 57 {
				return false
			}
		}
	}
	return true
}
