package store

import (
	mdl "final_project/models"
)

type AuthorStore interface {
	CreateAuthor(author mdl.Author) (mdl.Author, error)
	GetAuthor(id int) (mdl.Author, error)
	UpdateAuthor(id int, author mdl.Author) (mdl.Author, error)
	DeleteAuthor(id int) error
	GetAllAuthors() ([]mdl.Author, error)
}
