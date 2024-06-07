package bookstore

import (
	"errors"
	"sort"
)

type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book, n int) (Book, error) {
	if b.Copies <= 0 {
		return Book{}, errors.New("no copies left")
	}

	b.Copies -= n
	return b, nil
}

func GetAllBooks(catalog map[int64]Book) []Book {

	books := make([]Book, 0, len(catalog))

	for _, v := range catalog {
		books = append(books, v)
	}

	sort.Slice(books, func(i, j int) bool {
		return books[i].Title < books[j].Title
	})

	return books
}

func GetBook(catalog map[int64]Book, i int64) (Book, error) {

	b, ok := catalog[i]

	if !ok {
		return Book{}, errors.New("book not found")
	}

	return b, nil
}
