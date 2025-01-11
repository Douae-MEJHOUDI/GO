package store

type Stores struct {
	Books     BookStore
	Authors   AuthorStore
	Customers CustomerStore
	Orders    OrderStore
}

func NewStores() *Stores {
	return &Stores{
		Books:     NewBookStore(),
		Authors:   NewAuthorStore(),
		Customers: NewCustomerStore(),
		Orders:    NewOrderStore(),
	}
}
