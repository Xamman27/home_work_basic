package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

type CompareMode int

const (
	ByYear CompareMode = iota
	BySize
	ByRate
)

type BookComparator struct {
	mode CompareMode
}

func (b *Book) SetId(id int)            { b.id = id }
func (b *Book) GetId() int              { return b.id }
func (b *Book) SetTitle(title string)   { b.title = title }
func (b *Book) GetTitle() string        { return b.title }
func (b *Book) SetAuthor(author string) { b.author = author }
func (b *Book) GetAuthor() string       { return b.author }
func (b *Book) SetYear(year int)        { b.year = year }
func (b *Book) GetYear() int            { return b.year }
func (b *Book) SetSize(size int)        { b.size = size }
func (b *Book) GetSize() int            { return b.size }
func (b *Book) SetRate(rate float64)    { b.rate = rate }
func (b *Book) GetRate() float64        { return b.rate }

func NewComparator(mode CompareMode) *BookComparator {
	return &BookComparator{mode: mode}
}

func (c BookComparator) Compare(a, b Book) bool {
	switch c.mode {
	case ByYear:
		return a.year > b.year
	case BySize:
		return a.size > b.size
	case ByRate:
		return a.rate > b.rate
	default:
		return false
	}
}
func main() {
	b2 := Book{}
	b1 := Book{}
	b1.SetId(1)
	b1.SetTitle("Колобок")
	b1.SetAuthor("Vel")
	b1.SetYear(1999)
	b1.SetSize(50)
	b1.SetRate(1.5)

	b2.SetId(2)
	b2.SetTitle("Красная шапочка")
	b2.SetAuthor("Vel")
	b2.SetYear(20)
	b2.SetSize(60)
	b2.SetRate(5)

	cmp := NewComparator(ByYear)
	if cmp.Compare(b1, b2) {
		fmt.Println("b1 новее b2")
	} else {
		fmt.Println("b2 новее b1")
	}
}
