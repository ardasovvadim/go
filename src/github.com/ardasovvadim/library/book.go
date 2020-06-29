package main

import "fmt"

type Book struct {
	name   string
	papers int
}

func NewBook(name string, papers int) *Book {
	return &Book{
		name:   name,
		papers: papers,
	}
}

func (this *Book) String() string {
	return fmt.Sprintf("Book: {name = \"%s\", papers: %d}", this.name, this.papers)
}

func (this *Book) Equauls(book *Book) bool {
	return this.name == book.name && this.papers == book.papers
}

func RemoveBook(books []*Book, index int) []*Book {
	newArray := make([]*Book, len(books)-1)
	newArray = append(books[0:index], books[index+1:]...)
	return newArray
}
