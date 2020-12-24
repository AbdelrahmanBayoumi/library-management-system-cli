package main

import (
	"fmt"
	bookModel "library-management-system-cli/books"
	readerModel "library-management-system-cli/readers"
)

func main() {
	newBook := bookModel.Book{"3", "title_STR", "publicationDate_STR", "author_STR", "genre_STR", "publisher_STR", "language_STR"}

	fmt.Println("New book:", newBook.ToString())
	fmt.Println("==========================")
	fmt.Println(readerModel.GetAllReaders())

}
