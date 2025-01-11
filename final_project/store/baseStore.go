package store

type Stores struct {
	Books     BookStore
	Authors   AuthorStore
	Customers CustomerStore
	Orders    OrderStore
}

func NewStores() *Stores {
	stores := &Stores{}
	authorStore := NewAuthorStore(nil)
	bookStore := NewBookStore(authorStore)

	authorStore.books = bookStore

	stores.Authors = authorStore
	stores.Books = bookStore
	stores.Customers = NewCustomerStore()
	stores.Orders = NewOrderStore()
	return stores
}
