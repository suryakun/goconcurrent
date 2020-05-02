package main

import "fmt"

type Book struct {
	ID    int
	Title string
}

func (b Book) String() string {
	return fmt.Sprintf("ID %v \n Title: %q", b.ID, b.Title)
}

var books = []Book{
	Book{
		ID:    1,
		Title: "Test one",
	},
	Book{
		ID:    2,
		Title: "Test two",
	},
	Book{
		ID:    3,
		Title: "Test three",
	},
	Book{
		ID:    4,
		Title: "Test four",
	},
	Book{
		ID:    5,
		Title: "Test five",
	},
}
