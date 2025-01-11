package models

import (
	"errors"
)

var (
	ErrAuthorNotFound    = errors.New("Author not found")
	ErrFirstNameRequired = errors.New("first_name is required")
	ErrLastNAmeRequired  = errors.New("last_name is required")
	ErrBioRequired       = errors.New("bio is required")
)

type Author struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
}

func (a *Author) Validate() error {
	if a.FirstName == "" {
		return ErrFirstNameRequired
	}
	if a.LastName == "" {
		return ErrLastNAmeRequired
	}
	if a.Bio == "" {
		return ErrBioRequired
	}
	return nil
}
