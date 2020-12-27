package main

import (
	"fmt"
	bookRouter "library-management-system-cli/routers/books"
	readerRouter "library-management-system-cli/routers/readers"
	"net/http"
)

// ServeHTTP ...
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`library management system`))
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

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.HandleFunc("/books", bookRouter.ServeHTTP)
	http.HandleFunc("/books/search", bookRouter.SearchHandle)
	http.HandleFunc("/readers", readerRouter.ServeHTTP)
	http.HandleFunc("/readers/search", readerRouter.SearchHandle)
	err := http.ListenAndServe(":8050", nil)
	if err != nil {
		fmt.Println("FIRST SEVER IS DOWN!")
		err := http.ListenAndServe(":8051", nil)
		if err != nil {
			fmt.Println("SECOND SEVER IS DOWN!")
			panic(err)
		}
	}
}
