package main

import (
	"fmt"

	"github.com/fixme_my_friend/hw02_fix_app/printer"
	"github.com/fixme_my_friend/hw02_fix_app/reader"
)

func main() {
	var path string
	fmt.Println("Enter your path (default: data.json): ")
	_, err := fmt.Scanln(&path)
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println("Ошибка ввода:", err)
	}
	if path == "" {
		path = "data.json"
	}
	staff, err := reader.ReadJSON(path, -1)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	printer.PrintStaff(staff)
}
