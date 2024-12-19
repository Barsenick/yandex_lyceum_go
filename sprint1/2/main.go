package main

type Book struct {
	Title  string
	Author string
	Year   int
	Genre  string
}

func (b Book) NewBook(title, author string, year int, genre string) *Book {
	newBook := &Book{title, author, year, genre}
	return newBook
}
