package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	bookModel "library-management-system-cli/model/books"
	"net/http"
)

// ServeHTTP ...
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keys := r.URL.Query()
		sortMethod := keys.Get("sortBy")
		if sortMethod == "" {
			getAllBooks(w, r)
		} else {
			getBooksSorted(w, r, sortMethod)
		}
	case "POST":
		addBook(w, r)
	case "PUT":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	case "DELETE":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

// SearchHandle ...
func SearchHandle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keys := r.URL.Query()
		id := keys.Get("id")
		title := keys.Get("title")
		if id == "" && title == "" {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
			return
		} else if id != "" {
			// search by id and return book
			w.Header().Set("Content-Type", "application/json")
			book, flag := bookModel.GetAllBooks().GetBookByID(id)
			if flag != 1 {
				// no id found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"message": "no book with this id"}`))
			} else {
				// id found
				json, _ := json.Marshal(book)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(json))
			}
		} else if title != "" {
			// search by title and return book
			w.Header().Set("Content-Type", "application/json")
			book, flag := bookModel.GetAllBooks().GetBookByTitle(title)
			if flag != 1 {
				// no id found
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"message": "no book with this title"}`))
			} else {
				// id found
				json, _ := json.Marshal(book)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(json))
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message": "not found"}`))
		}
	case "POST":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	case "PUT":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	case "DELETE":
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "no body found"}`))
		return
	}
	var book bookModel.Book
	err := json.Unmarshal(body, &book)
	if err != nil {
		panic(err)
	}
	books := bookModel.GetAllBooks()
	book.InsertBook(&(books))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Book Added Successfully!"}`))
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllBooks() ...")
	w.Header().Set("Content-Type", "application/json")
	books := bookModel.GetAllBooks()
	json, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}

func getBooksSorted(w http.ResponseWriter, r *http.Request, sortMethod string) {
	fmt.Println("GetBooksSorted() ...")
	w.Header().Set("Content-Type", "application/json")
	books := bookModel.GetAllBooks()
	fmt.Println("sortMethod:", sortMethod)
	switch sortMethod {
	case "title":
		books.SortByTitle()
	case "publication date":
		books.SortByPublicationDate()
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
		return
	}
	json, _ := json.Marshal(books)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(json))
}
