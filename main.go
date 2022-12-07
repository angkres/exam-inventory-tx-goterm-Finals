package main

import (
	"exam-inventory/database"
	"exam-inventory/template"
)

func main() {
	db := database.GetConnection()
	template.Menu(db)
}
