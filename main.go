package main

import (
	"fmt"
	bookslib "library-management-system-cli/books"
)

func main() {
	newBook := bookslib.Book{"3", "title_STR", "publicationDate_STR", "author_STR", "genre_STR", "publisher_STR", "language_STR"}
	fmt.Println("New book:", newBook.ToString())
}
