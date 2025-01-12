package models

import (
	"errors"
	"time"
)

var (
	ErrCustomerNotFound         = errors.New("customer not found")
	ErrNameRequired             = errors.New("name is required")
	ErrEmailRequired            = errors.New("email is required")
	ErrStreetRequired           = errors.New("street is required")
	ErrCityRequired             = errors.New("city is required")
	ErrStateRequired            = errors.New("state is required")
	ErrPostalCodeRequired       = errors.New("postal_code is required")
	ErrCountryRequired          = errors.New("country is required")
	ErrCustomerNotSavedInMemory = errors.New("customer changes were not saved into memory")
	ErrCustomerHasOrders        = errors.New("can't delete cutomer with existing orders")
)

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   Address   `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return ErrNameRequired
	}
	if c.Email == "" {
		return ErrEmailRequired
	}
	if c.Address.Street == "" {
		return ErrStreetRequired
	}
	if c.Address.City == "" {
		return ErrCityRequired
	}
	if c.Address.State == "" {
		return ErrStateRequired
	}
	if c.Address.PostalCode == "" {
		return ErrPostalCodeRequired
	}
	if c.Address.Country == "" {
		return ErrCountryRequired
	}
	return nil
}
