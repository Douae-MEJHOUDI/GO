package store

import (
	mdl "final_project/models"
	"sync"
)

type InMemoryAuthorStore struct {
	mu      sync.RWMutex
	authors []mdl.Author
	nextID  int
}

func NewAuthorStore() *InMemoryAuthorStore {
	return &InMemoryAuthorStore{
		authors: []mdl.Author{},
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

	//author.ID = id
	err := author.Validate()
	if err != nil {
		return mdl.Author{}, err
	}
	for i, a := range s.authors {
		if a.ID == id {
			author.ID = id
			s.authors[i] = author
			return author, nil
		}
	}

	return mdl.Author{}, mdl.ErrAuthorNotFound
}

func (s *InMemoryAuthorStore) DeleteAuthor(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	//delete(s.authors, id)

	for i, author := range s.authors {
		if author.ID == id {
			s.authors = append(s.authors[:i], s.authors[i+1:]...)
			return nil
		}
	}

	return mdl.ErrAuthorNotFound
}

func (s *InMemoryAuthorStore) GetAllAuthors() ([]mdl.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.authors, nil
}
