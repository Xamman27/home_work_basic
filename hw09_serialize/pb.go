package main

import (
	"hw09_serialize/bookpb"

	"google.golang.org/protobuf/proto"
)

func (b *Book) ToProto() *bookpb.Book {
	return &bookpb.Book{
		Id:     int32(b.ID),
		Title:  b.Title,
		Author: b.Author,
		Year:   int32(b.Year),
		Size:   int32(b.Size),
		Rate:   b.Rate,
	}
}

func FromProto(pb *bookpb.Book) Book {
	return Book{
		ID:     int(pb.Id),
		Title:  pb.Title,
		Author: pb.Author,
		Year:   int(pb.Year),
		Size:   int(pb.Size),
		Rate:   pb.Rate,
	}
}

func SerializeBooksToProto(books []Book) ([]byte, error) {
	pbBooks := make([]*bookpb.Book, len(books))
	for i, b := range books {
		pbBooks[i] = b.ToProto()
	}
	bookList := &bookpb.BookList{Books: pbBooks}
	return proto.Marshal(bookList)
}

func DeserializeBooksFromProto(data []byte) ([]Book, error) {
	var bookList bookpb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, err
	}
	books := make([]Book, len(bookList.Books))
	for i, pb := range bookList.Books {
		books[i] = FromProto(pb)
	}
	return books, nil
}
