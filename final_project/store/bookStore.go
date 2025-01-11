package store

import (
	mdl "final_project/models"
)

type BookStore interface {
	CreateBook(book mdl.Book) (mdl.Book, error)
	GetBook(id int) (mdl.Book, error)
	UpdateBook(id int, book mdl.Book) (mdl.Book, error)
	DeleteBook(id int) error
	SearchBooks(criteria mdl.SearchCriteria) ([]mdl.Book, error)
	GetAllBooks() []mdl.Book
}
