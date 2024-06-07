package bookstore_test

import (
	"calculator/bookstore"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var catalog = map[int64]bookstore.Book{
	1: {
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		Copies: 3,
	},
	2: {
		Title:  "The Go Programming Language",
		Author: "Alan Donovan",
		Copies: 5,
	},
	3: {
		Title:  "Design Patterns",
		Author: "Erich Gamma",
		Copies: 2,
	},
}

func TestBookType(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		Copies: 3,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		Copies: 3,
	}

	result, error := bookstore.Buy(book, 2)

	if error != nil {
		t.Fatalf("Error: %v", error)
	}

	got := result.Copies

	if got != book.Copies-2 {
		t.Errorf("Want %v, got %v", book.Copies-2, got)
	}

}

func TestBuyInvalid(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title:  "The Art of Computer Programming",
		Author: "Donald Knuth",
		Copies: 0,
	}

	_, error := bookstore.Buy(book, 1)

	if error == nil {
		t.Errorf("Want an error, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	want := catalog

	got := bookstore.GetAllBooks(want)

	if !cmp.Equal(want[3], got[0]) {
		t.Errorf(cmp.Diff(want[1], got[0]))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()

	got, error := bookstore.GetBook(catalog, 1)

	if error != nil {
		t.Fatalf("Error: %v", error)
	}

	if !cmp.Equal(catalog[1], got) {
		t.Errorf(cmp.Diff(catalog[1], got))
	}
}

func TestGetBookNotFound(t *testing.T) {
	t.Parallel()

	_, err := bookstore.GetBook(catalog, 4)

	if err == nil {
		t.Errorf("Want an not found error, got nil")
	}
}
