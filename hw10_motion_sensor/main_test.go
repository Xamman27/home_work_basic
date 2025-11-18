package main

import (
	"testing"
	"time"
)

// Тестируем RandomNum: число в диапазоне [0, 99]
func TestRandomNum(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := RandomNum()
		if n < 0 || n >= 100 {
			t.Errorf("RandomNum() = %d, want in [0,99]", n)
		}
	}
}

// Тестируем GenerateNum: должно быть 10 чисел
func TestGenerateNum(t *testing.T) {
	ch := make(chan int)
	go GenerateNum(ch)

	nums := []int{}
	timeout := time.After(2 * time.Second)
loop:
	for {
		select {
		case n, ok := <-ch:
			if !ok {
				break loop
			}
			nums = append(nums, n)
		case <-timeout:
			t.Fatal("timeout waiting for GenerateNum")
		}
	}
	if len(nums) != 10 {
		t.Errorf("GenerateNum produced %d numbers, want 10", len(nums))
	}
}

// Тестируем Average: правильное среднее
func TestAverage(t *testing.T) {
	ch := make(chan int)
	result := make(chan float64)

	// Передаём 1, 2, 3, 4, 5 => среднее = 3.0
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	go Average(ch, result)

	select {
	case avg, ok := <-result:
		if !ok {
			t.Fatal("result channel closed unexpectedly")
		}
		if avg != 3.0 {
			t.Errorf("Average = %v, want 3.0", avg)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timeout waiting for Average")
	}
}
