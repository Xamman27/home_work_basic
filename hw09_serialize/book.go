package main

import (
	"encoding/json"
)

type Book struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Year   int     `json:"year"`
	Size   int     `json:"size"`
	Rate   float64 `json:"rate"`
}

func (b Book) MarshalJSON() ([]byte, error) {
	type Alias Book
	return json.Marshal(Alias(b))
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type Alias Book
	aux := &Alias{}
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}
	*b = Book(*aux)
	return nil
}

func SerializeBooksToJSON(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

func DeserializeBooksFromJSON(data []byte) ([]Book, error) {
	var books []Book
	err := json.Unmarshal(data, &books)
	return books, err
}
