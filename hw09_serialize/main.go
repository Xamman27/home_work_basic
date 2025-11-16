package main

import (
	"fmt"
	"log"
)

func main() {
	books := []Book{
		{ID: 1, Title: "Go", Author: "Alan", Year: 2020, Size: 300, Rate: 4.5},
		{ID: 2, Title: "Python", Author: "Guido", Year: 2015, Size: 250, Rate: 4.7},
	}

	// JSON
	jsonData, err := SerializeBooksToJSON(books)
	if err != nil {
		log.Fatalf("JSON Marshal error: %v", err)
	}
	fmt.Println("JSON:", string(jsonData))

	booksFromJSON, err := DeserializeBooksFromJSON(jsonData)
	if err != nil {
		log.Fatalf("JSON Unmarshal error: %v", err)
	}
	fmt.Printf("Books from JSON: %+v\n", booksFromJSON)

	// Proto
	protoData, err := SerializeBooksToProto(books)
	if err != nil {
		log.Fatalf("Proto Marshal error: %v", err)
	}
	fmt.Printf("Proto bytes: %v\n", protoData)

	booksFromProto, err := DeserializeBooksFromProto(protoData)
	if err != nil {
		log.Fatalf("Proto Unmarshal error: %v", err)
	}
	fmt.Printf("Books from Proto: %+v\n", booksFromProto)
}
