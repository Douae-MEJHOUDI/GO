package store

import (
	mdl "final_project/models"
	"strings"
	"sync"
)

type InMemoryBookStore struct {
	mu     sync.RWMutex
	books  []mdl.Book
	nextID int
}

func NewBookStore() *InMemoryBookStore {
	return &InMemoryBookStore{
		books:  []mdl.Book{},
		nextID: 1,
	}
}

func (s *InMemoryBookStore) CreateBook(book mdl.Book) (mdl.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := book.Validate()

	if err != nil {
		return mdl.Book{}, err
	}

	book.ID = s.nextID
	s.books = append(s.books, book)
	s.nextID++

	return book, nil
}

func (s *InMemoryBookStore) SearchBooks(criteria mdl.SearchCriteria) ([]mdl.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var results []mdl.Book

	if criteria.IsEmpty() {
		/*
			for _, book := range s.books {
				results = append(results, book)
			}*/
		return s.books, nil
	}

	for _, book := range s.books {
		if s.matchesCriteria(book, criteria) {
			results = append(results, book)
		}
	}

	return results, nil
}

func (s *InMemoryBookStore) matchesCriteria(book mdl.Book, criteria mdl.SearchCriteria) bool {
	if criteria.Title != "" && !strings.Contains(
		strings.ToLower(book.Title),
		strings.ToLower(criteria.Title)) {
		return false
	}
	/*
		if criteria.Author.FirstName != "" || criteria.Author.LastName != "" {
			authorName := strings.ToLower(book.Author.FirstName + " " + book.Author.LastName)
			if !strings.Contains(authorName, strings.ToLower(criteria.Author)) {
				return false
			}
		}*/

	return true
}

func (s *InMemoryBookStore) GetBook(id int) (mdl.Book, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, book := range s.books {
		if book.ID == id {
			return book, nil
		}
	}

	return mdl.Book{}, mdl.ErrBookNotFound
}

func (s *InMemoryBookStore) UpdateBook(id int, book mdl.Book) (mdl.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := book.Validate()
	if err != nil {
		return mdl.Book{}, err
	}

	for i, b := range s.books {
		if b.ID == id {
			book.ID = id
			s.books[i] = book
			return book, nil

		}
	}

	return mdl.Book{}, mdl.ErrBookNotFound

}

func (s *InMemoryBookStore) DeleteBook(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}

	return mdl.ErrBookNotFound
}
