package models

import (
	"errors"
	"time"
)

var (
	ErrEmptyTitle     = errors.New("book title cannot be empty")
	ErrInvalidPrice   = errors.New("book price must be greater than 0")
	ErrInvalidStock   = errors.New("book stock cannot be negative")
	ErrMissingAuthor  = errors.New("book must have an author")
	ErrEmptyGenres    = errors.New("book must have at least one genre")
	ErrInvalidPubDate = errors.New("publication date cannot be in the future")
	ErrBookNotFound   = errors.New("book not found")
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      Author    `json:"author"`
	Genres      []string  `json:"genres"`
	PublishedAt time.Time `json:"published_at"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return ErrEmptyTitle
	}
	if b.Price <= 0 {
		return ErrInvalidPrice
	}
	if b.Stock < 0 {
		return ErrInvalidStock
	}
	if b.Author.ID == 0 {
		return ErrMissingAuthor
	}
	if len(b.Genres) == 0 {
		return ErrEmptyGenres
	}
	if b.PublishedAt.After(time.Now()) {
		return ErrInvalidPubDate
	}
	return nil
}
