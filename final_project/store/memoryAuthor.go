package store

import (
	mdl "final_project/models"
	"sync"
)

type InMemoryAuthorStore struct {
	mu      sync.RWMutex
	authors []mdl.Author
	books   BookStore
	nextID  int
}

func NewAuthorStore(books BookStore) *InMemoryAuthorStore {
	return &InMemoryAuthorStore{
		authors: []mdl.Author{},
		books:   books,
		nextID:  1,
	}
}

func (s *InMemoryAuthorStore) CreateAuthor(author mdl.Author) (mdl.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := author.Validate()
	if err != nil {
		return mdl.Author{}, err
	}
	author.ID = s.nextID
	s.authors = append(s.authors, author)
	s.nextID++
	return author, nil
}

func (s *InMemoryAuthorStore) GetAuthor(id int) (mdl.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, author := range s.authors {
		if author.ID == id {
			return author, nil
		}
	}

	return mdl.Author{}, mdl.ErrAuthorNotFound

}

func (s *InMemoryAuthorStore) UpdateAuthor(id int, author mdl.Author) (mdl.Author, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := author.Validate()
	if err != nil {
		return mdl.Author{}, err
	}
	index := -1
	for i, a := range s.authors {
		if a.ID == id {
			index = i
		}
	}

	if index == -1 {
		return mdl.Author{}, mdl.ErrAuthorNotFound
	}

	author.ID = id
	s.authors[index] = author
	for _, book := range s.books.GetAllBooks() {
		if book.Author.ID == id {
			book.Author = author
		}
	}

	return author, nil
}

func (s *InMemoryAuthorStore) DeleteAuthor(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	authIndex := -1

	for i, author := range s.authors {
		if author.ID == id {
			authIndex = i
			break
		}
	}

	if authIndex == -1 {
		return mdl.ErrAuthorNotFound
	}

	for _, book := range s.books.GetAllBooks() {
		if book.Author.ID == id {
			return mdl.ErrAuthorHasBooks
		}
	}

	s.authors = append(s.authors[:authIndex], s.authors[authIndex+1:]...)
	return nil
}

func (s *InMemoryAuthorStore) GetAllAuthors() ([]mdl.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.authors, nil
}
