package main

import (
	"errors"
	"fmt"
	"math"
)

// Интерфейс Shape с методом Area
type Shape interface {
	Area() float64
}

// Структура Круг
type Circle struct {
	Radius float64
}

// Реализация метода Area для круга
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Структура Прямоугольник
type Rectangle struct {
	Width  float64
	Height float64
}

// Реализация метода Area для прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Структура Треугольник
type Triangle struct {
	Base   float64
	Height float64
}

// Реализация метода Area для треугольника
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Функция для подсчета площади
func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("переданный объект не является фигурой")
	}
	return shape.Area(), nil
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}
	t := Triangle{Base: 8, Height: 6}
	notShape := 123 // просто число, не фигура

	// Круг
	area, err := calculateArea(c)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Круг: радиус %.0f Площадь: %v\n", c.Radius, area)
	}

	// Прямоугольник
	area, err = calculateArea(r)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Прямоугольник: ширина %.0f, высота %.0f Площадь: %v\n", r.Width, r.Height, area)
	}

	// Треугольник
	area, err = calculateArea(t)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Треугольник: основание %.0f, высота %.0f Площадь: %v\n", t.Base, t.Height, area)
	}

	// Не фигура
	area, err = calculateArea(notShape)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Площадь: %v\n", area)
	}
}
