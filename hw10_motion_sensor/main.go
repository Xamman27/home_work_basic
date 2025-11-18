package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomNum() int {
	return rng.Intn(100)
}

func GenerateNum(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- RandomNum()
	}
	close(ch)
}

func Average(ch chan int, result chan float64) {
	sum := 0
	count := 0
	for num := range ch {
		sum += num
		count++
	}
	if count > 0 {
		result <- float64(sum) / float64(count)
	} else {
		result <- 0
	}
	close(result)
}

func main() {
	ch := make(chan int)
	result := make(chan float64)
	go GenerateNum(ch)
	go Average(ch, result)

	avg := <-result
	fmt.Printf("Среднее: %.2f\n", avg)
}
