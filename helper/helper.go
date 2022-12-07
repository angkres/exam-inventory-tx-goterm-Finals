package helper

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/alok87/goutils/pkg/random"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Fungsi untuk kembali ke menu
func BackHandler() {
	fmt.Print("Tekan enter untuk kembali ke menu")
	var back int
	fmt.Scanln(&back)
}

func PauseHandler() {
	var jeda string
	fmt.Print("Tekan enter untuk melanjutkan")
	fmt.Scanln(&jeda)
	ClearScreen()
}

func DateToString(date time.Time) string {
	dateString := date.Format("2006-01-02")
	return dateString
}

func SplitString(orderNumber string) (string, string) {
	orderNumberSlice := strings.Split(orderNumber, "-")
	typeOrder := orderNumberSlice[0]
	number := orderNumberSlice[1]
	return typeOrder, number
}

func RandomNumber() string {
	random := random.RangeInt(10000, 99999, 1)
	number := strconv.Itoa(random[0])
	return number
}

func NullStringtoStringEmail(email *sql.NullString) string {
	var emailString string
	if email.Valid {
		emailString = email.String
		return emailString
	}
	return ""
}

func NullStringtoStringPhone(phone *sql.NullString) string {
	var phoneString string
	if phone.Valid {
		phoneString = phone.String
		return phoneString
	}
	return ""
}
