package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Reader struct {
	name            string
	booksForReading []*Book
	booksForGetting []*Book
	channel         chan string
}

func NewReader(name string) *Reader {
	return &Reader{
		name:            name,
		booksForReading: make([]*Book, 0, 10),
		booksForGetting: make([]*Book, 0, 10),
		channel:         make(chan string),
	}
}

func (this *Reader) String() string {
	return fmt.Sprintf("Reader {name = \"%s\", books: %d}", this.name, len(this.booksForGetting))
}

func (this *Reader) GetWantedBooks(books []*Book) ([]*Book, []*Book) {
	// [0-1]
	amountForReading := rand.Intn(2)
	// [1-2]
	amountForGetting := rand.Intn(2) + 1
	forReading := make([]*Book, 0, 1)
	forGetting := make([]*Book, 0, 2)

	if amountForReading > 0 {
		bookIndex := rand.Intn(len(books))
		forReading = append(forReading, books[bookIndex])
		books = RemoveBook(books, bookIndex)
	}

	{
		i := 0
		for i < amountForGetting {
			bookIndex := rand.Intn(len(books))
			forGetting = append(forGetting, books[bookIndex])
			books = RemoveBook(books, bookIndex)
			i++
		}
	}

	return forReading, forGetting
}

func (this *Reader) Run() {
	log(this.name, "Start working...")
	time.Sleep(time.Second)

	log(this.name, "Going to the library")
	time.Sleep(time.Second)

	log(this.name, "Going to the reception")
	time.Sleep(time.Second)

	Reception <- this
	answer := <-this.channel

	if answer == "Processed" && len(this.booksForReading) > 0 {
		log(this.name, "Going to the reading room")
		time.Sleep(time.Second)

		ReadingRoom <- this
		answer = <-this.channel

		log(this.name, "Going to the reception")
		time.Sleep(time.Second)

		log(this.name, "Put the books")
		this.PushBook(Librarians)
		time.Sleep(time.Second * 2)
	}

	log(this.name, "Left the library")
}

func (this *Reader) ReadBook() {
	for _, b := range this.booksForReading {
		log(this.name, "Start reading book "+b.name)
		time.Sleep(time.Second * 5)
		log(this.name, "End reading book "+b.name)
	}
}

func (this *Reader) PushBook(librarians *Librarian) {
	librarians.books = append(librarians.books, this.booksForReading[0])
	this.booksForReading = make([]*Book, 0, 10)
}
