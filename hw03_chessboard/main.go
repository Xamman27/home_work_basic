package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Welcome to Chessboard ! \nEnter size: ")
	n, err := fmt.Scanf("%d", &size)
	if n == 1 && err == nil {
		fmt.Println("Size:", size)
	} else {
		size = 8
		fmt.Println("Invalid input. Default size 8 is used.")
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i%2 == 0 && j%2 == 0) || (i%2 == 1 && j%2 == 1) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
