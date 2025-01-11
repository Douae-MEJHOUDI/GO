package store

import (
	mdl "final_project/models"
	"strings"
	"sync"
)

type InMemoryBookStore struct {
	mu      sync.RWMutex
	books   []mdl.Book
	authors AuthorStore
	nextID  int
}

func NewBookStore(authors AuthorStore) *InMemoryBookStore {
	return &InMemoryBookStore{
		books:   []mdl.Book{},
		authors: authors,
		nextID:  1,
	}
}

func (s *InMemoryBookStore) CreateBook(book mdl.Book) (mdl.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := book.Validate()

	if err != nil {
		return mdl.Book{}, err
	}

	realAuthor, err := s.authors.GetAuthor(book.Author.ID)
	if err != nil {
		return mdl.Book{}, err
	}

	book.Author = realAuthor

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
		return s.books, nil
	}

	for _, book := range s.books {
		matches := true

		if criteria.Title != "" {
			matches = matches && strings.Contains(
				strings.ToLower(book.Title),
				strings.ToLower(criteria.Title))
		}
		if criteria.Author != "" {
			authorName := strings.ToLower(book.Author.FirstName + " " + book.Author.LastName)
			matches = matches && strings.Contains(
				authorName,
				strings.ToLower(criteria.Author),
			)
		}

		if criteria.MinPrice != nil {
			matches = matches && book.Price >= *criteria.MinPrice
		}

		if criteria.MaxPrice != nil {
			matches = matches && book.Price <= *criteria.MaxPrice
		}

		if len(criteria.Genres) > 0 {
			exists := false

			for _, genre := range criteria.Genres {
				for _, bookGenre := range book.Genres {
					if genre == bookGenre {
						exists = true
						break
					}
				}
				if exists {
					break
				}
			}
			matches = matches && exists
		}

		if criteria.InStock != nil {
			instock := book.Stock > 0
			matches = matches && instock == *criteria.InStock
		}

		if matches {
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

	realAuthor, err := s.authors.GetAuthor(book.Author.ID)
	if err != nil {
		return mdl.Book{}, err
	}

	book.Author = realAuthor

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

func (s *InMemoryBookStore) GetAllBooks() []mdl.Book {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.books
}
