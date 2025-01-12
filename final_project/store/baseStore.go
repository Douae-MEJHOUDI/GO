package store

import "errors"

type Stores struct {
	Books     BookStore
	Authors   AuthorStore
	Customers CustomerStore
	Orders    OrderStore
}

func NewStores() (*Stores, error) {
	data, err := NewDataManager("./data")
	if err != nil {
		return nil, errors.New("Problem in fetching data: " + err.Error())
	}

	stores := &Stores{}
	authorStore := NewAuthorStore(nil, data)
	bookStore := NewBookStore(authorStore, data)

	authorStore.books = bookStore

	stores.Authors = authorStore
	stores.Books = bookStore
	stores.Customers = NewCustomerStore()
	stores.Orders = NewOrderStore()
	return stores, nil
}
