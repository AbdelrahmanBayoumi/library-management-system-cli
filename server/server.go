package main

import (
	"fmt"
	bookRouter "library-management-system-cli/routers/books"
	readerRouter "library-management-system-cli/routers/readers"
	"net/http"
)

func main() {
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
