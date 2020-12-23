package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library-management-system-cli/utils"
	"log"
	"os"
)

// Books struct which contains
// an array of books
type Books struct {
	Books []Book `json:"books"`
}

// Book struct which contains the book data
type Book struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	PublicationDate string `json:"publicationDate"`
	Author          string `json:"author"`
	Genre           string `json:"genre"`
	Publisher       string `json:"publisher"`
	Language        string `json:"language"`
}

func main() {
	booksFileName := "books.json"
	err := utils.CheckFile(booksFileName)
	if err != nil {
		log.Println(err)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(booksFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened books.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Books array
	var books Books

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'books' which we defined above
	json.Unmarshal(byteValue, &books)

	// we iterate through every book within our books array and
	// print out the book data
	for i := 0; i < len(books.Books); i++ {
		fmt.Println("books id: " + books.Books[i].ID)
		fmt.Println("books Age: " + books.Books[i].Title)
	}

	newBook := Book{"3", "title_STR", "publicationDate_STR", "author_STR", "genre_STR", "publisher_STR", "language_STR"}
	fmt.Println("New book:", newBook)

	books.Books = append(books.Books, newBook)

	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(books)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(booksFileName, dataBytes, 0644)
	if err != nil {
		log.Println(err)
	}

	// we iterate through every book within our books array and
	// print out the book data
	for i := 0; i < len(books.Books); i++ {
		fmt.Println("books id: " + books.Books[i].ID)
		fmt.Println("books Age: " + books.Books[i].Title)
	}
}
