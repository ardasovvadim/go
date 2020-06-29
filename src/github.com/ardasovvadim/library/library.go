package main

import (
	"fmt"
	"time"
)

var Books = []*Book{
	NewBook("Book 1", 150),
	NewBook("Book 2", 250),
	NewBook("Book 3", 120),
	NewBook("Book 4", 135),
	NewBook("Book 5", 200),
	NewBook("Book 6", 220),
	NewBook("Book 7", 220),
}

var Librarians = NewLibrarian("Librarian 1", Books)

var Readers = []*Reader{
	NewReader("Reader 1"),
	NewReader("Reader 2"),
	NewReader("Reader 3"),
}

var Reception = make(chan *Reader)
var ReadingRoom = make(chan *Reader)

func log(name string, msg string) {
	t := time.Now().Format("15:04:05.000000")
	message := fmt.Sprintf("%s | %s: %s", t, name, msg)
	fmt.Println(message)
}

func main() {
	go ReadingRoomRun()
	go Librarians.Run()
	for _, r := range Readers {
		go r.Run()
		time.Sleep(time.Second)
	}
	fmt.Scanln()
	Reception <- nil
	ReadingRoom <- nil
}
