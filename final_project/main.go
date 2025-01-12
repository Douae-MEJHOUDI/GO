package main

import (
	"final_project/handlers"
	"final_project/service"
	"final_project/store"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")

	stores, err := store.NewStores()
	if err != nil {
		log.Println(err)
	}
	generator := service.NewGenerator(stores)

	go generator.GenerateReport()

	handler := handlers.NewHandler(stores)
	authorHandler := handlers.NewAuthorHandler(handler)
	bookHandler := handlers.NewBookHandler(handler)
	orderHandler := handlers.NewOrderHandler(handler)
	customerHandler := handlers.NewCustomerHandler(handler)

	http.HandleFunc("/books", bookHandler.BooksRequestHandler)
	http.HandleFunc("/books/", bookHandler.BookRequestHandler)

	http.HandleFunc("/authors", authorHandler.AuthorsRequestHandler)
	http.HandleFunc("/authors/", authorHandler.AuthorRequestHandler)

	http.HandleFunc("/orders", orderHandler.OrdersHandler)
	http.HandleFunc("/orders/", orderHandler.OrderHandler)

	http.HandleFunc("/customers", customerHandler.CustomersHandler)
	http.HandleFunc("/customers/", customerHandler.CustomerHandler)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)

}
