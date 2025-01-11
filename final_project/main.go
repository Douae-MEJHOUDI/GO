package main

import (
	"final_project/handlers"
	"final_project/service"
	"final_project/store"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	stores := store.NewStores()
	generator := service.NewGenerator(stores)

	go generator.GenerateReport()

	handler := handlers.NewHandler(stores)
	authorHandler := handlers.NewAuthorHandler(handler)
	bookHandler := handlers.NewBookHandler(handler)

	http.HandleFunc("/books", bookHandler.BooksRequestHandler)
	http.HandleFunc("/books/", bookHandler.BookRequestHandler)

	http.HandleFunc("/authors", authorHandler.AuthorsRequestHandler)
	http.HandleFunc("/authors/", authorHandler.AuthorRequestHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
