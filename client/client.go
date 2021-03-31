package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	booksModel "library-management-system-cli/model/books"
	readersModel "library-management-system-cli/model/readers"
	"library-management-system-cli/utils"
	"net/http"
)

// PORT for server
var PORT string

// HOST for server
const HOST = "http://localhost"

func testPort() {
	body, statusCode := getResponseGET(HOST + ":8050")
	if statusCode == 0 || string(body) != "library management system" {
		body, statusCode := getResponseGET(HOST + ":8051")
		if statusCode == 0 || string(body) != "library management system" {
			fmt.Println("ERROR IN CONNECTING TO SERVER!")
		} else {
			PORT = "8051"
		}
	} else {
		PORT = "8050"
	}
}

func getResponseGET(endpoint string) ([]byte, int) {
	response, err := http.Get(endpoint)
	if err != nil {
		return nil, 0
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	return body, response.StatusCode
}

func getBook(endpoint string) (int, booksModel.Book) {
	body, statusCode := getResponseGET(endpoint)
	if statusCode != 200 {
		return statusCode, booksModel.Book{}
	}
	var book booksModel.Book
	json.Unmarshal(body, &book)
	return 1, book
}

func getBooks(endpoint string) (int, booksModel.Books) {
	body, statusCode := getResponseGET(endpoint)
	if statusCode != 200 {
		return 0, booksModel.Books{}
	}
	var books booksModel.Books
	json.Unmarshal(body, &books)
	return 1, books
}

func getReader(endpoint string) (int, readersModel.Reader) {
	body, statusCode := getResponseGET(endpoint)
	if statusCode != 200 {
		return 0, readersModel.Reader{}
	}
	var reader readersModel.Reader
	json.Unmarshal(body, &reader)
	return 1, reader
}

func getReaders(endpoint string) (int, readersModel.Readers) {
	body, statusCode := getResponseGET(endpoint)
	if statusCode != 200 {
		return 0, readersModel.Readers{}
	}
	var readers readersModel.Readers
	json.Unmarshal(body, &readers)
	return 1, readers
}

func deleteReader(endpoint string) (int, readersModel.Reader) {
	req, err := http.NewRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return 0, readersModel.Reader{}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, readersModel.Reader{}
	}
	if resp.StatusCode != 200 {
		return 0, readersModel.Reader{}
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var reader readersModel.Reader
	json.Unmarshal(bodyBytes, &reader)
	return 1, reader
}

func main() {
	fmt.Println("-----------------------------------------")
	fmt.Println("-------- Library Management CLI ---------")
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("----------- Choose Operation ------------")
		fmt.Println("0) EXIT")
		fmt.Println("1) Books")
		fmt.Println("2) Local Reader")
		var flag int
		flag = handlePrompets()
		if flag == 0 {
			break
		}
	}
}

func handlePrompets() int {
	var choice string
	fmt.Scan(&choice)
	utils.ClearTerminal()
	switch choice {
	case "1":
		handleBooksPrompets()
	case "2":
		handleReadersPrompets()
	case "0":
		return 0
	}
	return 1
}

func handleBooksPrompets() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("--------- Choose Operation ----------")
		fmt.Println("-----------------------------------------")
		fmt.Println("1) Get all books")
		fmt.Println("2) Get all books sorted by title")
		fmt.Println("3) Get all books sorted by publication date")
		fmt.Println("4) Search by id")
		fmt.Println("5) Search by title")
		fmt.Println("6) Add book")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		utils.ClearTerminal()
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "2":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books?sortBy=title")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "3":
			testPort()
			flag, books := getBooks(HOST + ":" + PORT + "/books?sortBy=publication%20date")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				books.PrintAll()
			}
		case "4":
			fmt.Print("Enter book ID: ")
			var id string
			fmt.Scan(&id)
			testPort()
			flag, book := getBook(HOST + ":" + PORT + "/books/search?id=" + id)
			if flag != 1 {
				fmt.Println("There is No book with this ID")
			} else {
				fmt.Println("Book is:\n" + book.ToString() + "\n")
			}
		case "5":
			title := utils.ScanLine("Enter book title: ", "ERROR, Enter Book title: ")
			title = utils.ReplaceURLSpaces(title)
			testPort()
			flag, book := getBook(HOST + ":" + PORT + "/books/search?title=" + title)
			if flag != 1 {
				fmt.Println("There is No book with this title")
			} else {
				fmt.Println("Book is:\n" + book.ToString() + "\n")
			}
		case "6":
			book := scanBook()
			json, _ := json.Marshal(book)
			testPort()
			response, err := http.Post(HOST+":"+PORT+"/books", "application/json", bytes.NewBuffer([]byte(json)))
			if err != nil {
				fmt.Println("ERROR !")
			} else {
				if response.StatusCode != 200 {
					fmt.Println("ERROR !")
				} else {
					fmt.Println("Book added successfully!")
				}
			}
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}

func scanBook() booksModel.Book {
	fmt.Println("Enter book data:-")
	var book booksModel.Book
	book.ID = utils.ScanLine("Enter book ID: ", "ERROR, Enter Book ID: ")
	book.Title = utils.ScanLine("Enter book Title: ", "ERROR, Enter Book Title: ")
	book.PublicationDate = utils.ScanLine("Enter book PublicationDate: ", "ERROR, Enter Book PublicationDate: ")
	book.Author = utils.ScanLine("Enter book Author: ", "ERROR, Enter Book Author: ")
	book.Genre = utils.ScanLine("Enter book Genre: ", "ERROR, Enter Book Genre: ")
	book.Publisher = utils.ScanLine("Enter book Publisher: ", "ERROR, Enter Book Publisher: ")
	book.Language = utils.ScanLine("Enter book Language: ", "ERROR, Enter Book Language: ")
	return book
}

func handleReadersPrompets() {
	for {
		fmt.Println("-----------------------------------------")
		fmt.Println("--------- Choose Operation ----------")
		fmt.Println("1) Get all readers")
		fmt.Println("2) Search by id")
		fmt.Println("3) Search by name")
		fmt.Println("4) Add reader")
		fmt.Println("5) Remove reader")
		fmt.Println("0) BACK")
		var choice string
		fmt.Scan(&choice)
		utils.ClearTerminal()
		fmt.Println("-----------------------------------------")
		switch choice {
		case "1":
			testPort()
			flag, readers := getReaders(HOST + ":" + PORT + "/readers")
			if flag != 1 {
				fmt.Println("ERROR !")
			} else {
				readers.PrintAll()
			}
		case "2":
			fmt.Println("Enter local reader ID:")
			var id string
			fmt.Scan(&id)
			testPort()
			flag, reader := getReader(HOST + ":" + PORT + "/readers/search?id=" + id)
			if flag != 1 {
				fmt.Println("There is NO reader with this ID")
			} else {
				fmt.Println("Reader is:\n" + reader.ToString())
			}
		case "3":
			name := utils.ScanLine("Enter local reader name: ", "ERROR, Enter local reader name: ")
			name = utils.ReplaceURLSpaces(name)
			testPort()
			flag, reader := getReader(HOST + ":" + PORT + "/readers/search?name=" + name)
			if flag != 1 {
				fmt.Println("There is No reader with this name")
			} else {
				fmt.Println("Reader is:\n" + reader.ToString())
			}
		case "4":
			reader := scanReader()
			json, _ := json.Marshal(reader)
			testPort()
			response, err := http.Post(HOST+":"+PORT+"/readers", "application/json", bytes.NewBuffer([]byte(json)))
			if err != nil {
				fmt.Println("ERROR !")
			} else {
				if response.StatusCode != 200 {
					fmt.Println("ERROR !")
				} else {
					fmt.Println("Reader added successfully!")
				}
			}
		case "5":
			fmt.Println("Enter local reader ID:")
			var id string
			fmt.Scan(&id)
			testPort()
			flag, reader := deleteReader(HOST + ":" + PORT + "/readers?id=" + id)
			if flag != 1 {
				fmt.Println("There is NO reader with this ID")
			} else {
				fmt.Println("Reader Deleted Successfully!\n" + reader.ToString())
			}
		case "0":
			return
		default:
			fmt.Println("INPUT ERROR !")
		}
	}
}

func scanReader() readersModel.Reader {
	fmt.Println("Enter local reader data:-")
	var reader readersModel.Reader
	reader.ID = utils.ScanLine("Enter reader ID: ", "ERROR, Enter reader ID: ")
	reader.Name = utils.ScanLine("Enter reader Name: ", "ERROR, Enter reader Name: ")
	reader.Gender = utils.ScanLine("Enter reader Gender: ", "ERROR, Enter reader Gender: ")
	reader.Birthday = utils.ScanLine("Enter reader Birthday: ", "ERROR, Enter reader Birthday: ")
	reader.Height = utils.ScanLine("Enter reader Height: ", "ERROR, Enter reader Height: ")
	reader.Weight = utils.ScanLine("Enter reader Weight: ", "ERROR, Enter reader Weight: ")
	reader.Employment = utils.ScanLine("Enter reader Employment: ", "ERROR, Enter reader Employment: ")
	return reader
}
