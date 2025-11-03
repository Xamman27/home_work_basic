package main

import (
	"errors"
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

type Triangle struct {
	Base, Height float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, errors.New("Переданный объектт не является фигурой")
	}
	return shape.Area(), nil
}

func main() {
	c := Circle{Radius: 5}
	r := Rectangle{Width: 10, Height: 5}
	t := Triangle{Base: 8, Height: 6}
	notShape := 123

	area, err := CalculateArea(c)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Радиус круга %.0f, площадь %v\n", c.Radius, area)
	}

	area, err = CalculateArea(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ширина  прямоугольника %.0f, высота %.0f, площадь %v\n", r.Width, r.Height, area)
	}

	area, err = CalculateArea(t)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Основание  треугольника %.0f, высота %.0f, площадь %v\n", t.Height, t.Base, area)
	}

	area, err = CalculateArea(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ширина  прямоугольника %.0f, высота %.0f, площадь %v\n", r.Width, r.Height, area)
	}

	area, err = CalculateArea(notShape)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Площадь %v\n", area)
	}
}
