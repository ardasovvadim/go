package main

import (
	"fmt"
	"time"
)

type Librarian struct {
	name  string
	books []*Book
}

func NewLibrarian(name string, books []*Book) *Librarian {
	return &Librarian{
		name:  name,
		books: books,
	}
}

func (this *Librarian) String() string {
	return fmt.Sprintf("Librarian {name = \"%s\", books: %d}", this.name, len(this.books))
}

func (this *Librarian) IsExistsBook(book *Book) bool {
	for i := range this.books {
		if book.Equauls(this.books[i]) {
			return true
		}
	}
	return false
}

func (this *Librarian) TakeBook(book *Book) *Book {
	index := -1
	for i := range this.books {
		if this.books[i].Equauls(book) {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	} else {
		b := this.books[index]
		this.books = RemoveBook(this.books, index)
		return b
	}
}

func (this *Librarian) Run() {
	log(this.name, "Start working...")

	for true {
		reader := <-Reception

		log(this.name, "Start process "+reader.name)

		if reader == nil {
			break
		}

		copyBooks := make([]*Book, len(Books))
		copy(copyBooks, Books)
		forReading, forGetting := reader.GetWantedBooks(copyBooks)

		for _, b := range forReading {
			if this.IsExistsBook(b) {
				reader.booksForReading = append(reader.booksForReading, this.TakeBook(b))
				log(this.name, "Give "+b.name+" for reading")
				time.Sleep(time.Second)
			}
		}

		for _, b := range forGetting {
			if this.IsExistsBook(b) {
				reader.booksForGetting = append(reader.booksForGetting, this.TakeBook(b))
				log(this.name, "Give "+b.name+" for getting")
				time.Sleep(time.Second)
			}
		}

		log(this.name, "End process "+reader.name)

		reader.channel <- "Processed"
	}

	log(this.name, "Stopped")
}

func ReadingRoomRun() {
	name := "ReadingRoom"

	log(name, "Start working...")

	for true {
		reader := <-ReadingRoom
		log(reader.name, "Enter to the reading room")
		reader.ReadBook()
		log(reader.name, "Left the reading room")
		reader.channel <- "Finished"
	}

	log(name, "Stopped")
}
