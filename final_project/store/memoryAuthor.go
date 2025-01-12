package store

import (
	mdl "final_project/models"
	"log"
	"sync"
)

type InMemoryAuthorStore struct {
	mu      sync.RWMutex
	authors []mdl.Author
	books   BookStore
	nextID  int
	data    *DataManager
}

type DataMAuthor struct {
	Authors []mdl.Author `json:"authors`
	NextID  int          `json: next_id`
}

func NewAuthorStore(bookstore BookStore, data *DataManager) *InMemoryAuthorStore {
	store := &InMemoryAuthorStore{
		authors: []mdl.Author{},
		books:   bookstore,
		nextID:  1,
		data:    data,
	}

	var dataM DataMAuthor
	err := data.LoadData("authors.json", &dataM)
	if err != nil {
		log.Printf("Failed to load authors: %v\n", err)
	} else if len(dataM.Authors) > 0 {
		store.authors = dataM.Authors
		store.nextID = dataM.NextID
	}
	return store
}

func (s *InMemoryAuthorStore) saveAuthorData() error {
	var dataM DataMAuthor
	dataM.Authors = s.authors
	dataM.NextID = s.nextID

	return s.data.SaveData("authors.json", dataM)
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

	err = s.saveAuthorData()

	if err != nil {
		return mdl.Author{}, mdl.ErrAuthorNotSavedInMemory
	}

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
			err = s.saveAuthorData()
			if err != nil {
				return mdl.Author{}, mdl.ErrAuthorNotSavedInMemory
			}

			return author, nil
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
	err := s.saveAuthorData()

	if err != nil {
		return mdl.ErrAuthorNotSavedInMemory
	}

	return nil
}

func (s *InMemoryAuthorStore) GetAllAuthors() ([]mdl.Author, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.authors, nil
}
