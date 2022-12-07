package template

import (
	"exam-inventory/helper"
	"fmt"
)

func (template *orderTemplate) AddSOrder() {
	helper.ClearScreen()
	var nameCustomer, email, phone string
	var typeOrder = "SO"

	fmt.Println("Add Sales Order - Customer")
	fmt.Println("=============================")
	template.InputName(&nameCustomer)
	template.InputEmail(&email)
	template.InputPhone(&phone)

	helper.PauseHandler()

	template.InputOrders(&nameCustomer, &email, &phone, &typeOrder)
}
