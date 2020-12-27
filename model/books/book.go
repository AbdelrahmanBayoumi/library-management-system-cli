package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"library-management-system-cli/utils"
	"log"
	"os"
	"sort"
	"strings"
)

const fileName string = "books.json"

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

// PrintAll ...
func (books Books) PrintAll() {
	for i := 0; i < len(books.Books); i++ {
		fmt.Println(books.Books[i].ToString())
	}
}

// GetJSONstring ...
func GetJSONstring() string {
	err := utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return string(byteValue)
}

// GetAllBooks ...
func GetAllBooks() Books {
	err := utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// fmt.Println("byteValue:", string(byteValue))
	// we initialize our Books array
	var books Books

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'books' which we defined above
	json.Unmarshal(byteValue, &books)

	return books
}

// ToString ...
func (b Book) ToString() string {
	// fmt.Println(b)
	return "ID: " + b.ID + ", " + "Title: " + b.Title + ", " + "Publication Date: " + b.PublicationDate + ", " + "Author: " + b.Author + ", " + "Genre: " + b.Genre + ", " + "Publisher: " + b.Publisher + ", " + "Language: " + b.Language
}

// InsertBook function to insert a book books.json
func (b Book) InsertBook(books *Books) {
	(*books).Books = append((*books).Books, b)
	// Preparing the data to be marshalled and written.
	dataBytes, err := json.Marshal(*books)
	if err != nil {
		log.Println(err)
	}

	err = utils.CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

// GetBookByID ...
func (books Books) GetBookByID(id string) (Book, int) {
	for i := 0; i < len(books.Books); i++ {
		if books.Books[i].ID == id {
			return books.Books[i], 1
		}
	}
	return Book{}, 0
}

// GetBookByTitle ...
func (books Books) GetBookByTitle(title string) (Book, int) {
	for i := 0; i < len(books.Books); i++ {
		if strings.ToLower(books.Books[i].Title) == strings.ToLower(title) {
			return books.Books[i], 1
		}
	}
	return Book{}, 0
}

// SortByPublicationDate ...
func (books Books) SortByPublicationDate() {
	sort.SliceStable(books.Books, func(i, j int) bool {
		return books.Books[i].PublicationDate < books.Books[j].PublicationDate
	})
}

// SortByTitle ...
func (books Books) SortByTitle() {
	sort.SliceStable(books.Books, func(i, j int) bool {
		return books.Books[i].Title < books.Books[j].Title
	})
}

func main() {
	books := GetAllBooks()

	books.PrintAll()
	// fmt.Println("===================================")

	// book1, flag := books.GetBookByID("5")
	// if flag != 1 {
	// 	log.Println("ERROR in GetBookByID")
	// } else {
	// 	fmt.Println(book1.ToString())
	// }

	// fmt.Println("===================================")

	// book1, flag = books.GetBookByTitle("title_STR")
	// if flag != 1 {
	// 	log.Println("ERROR in GetBookByTitle")
	// } else {
	// 	fmt.Println(book1.ToString())
	// }

	// books.SortByTitle()
	fmt.Println("===================================")
	books.PrintAll()
}
