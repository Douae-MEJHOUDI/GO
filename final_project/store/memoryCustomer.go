package store

import (
	mdl "final_project/models"
	"sync"
)

type InMemoryCustomerStore struct {
	mu        sync.RWMutex
	customers []mdl.Customer
	nextID    int
}

func NewCustomerStore() *InMemoryCustomerStore {
	return &InMemoryCustomerStore{
		customers: []mdl.Customer{},
		nextID:    1,
	}
}

func (s *InMemoryCustomerStore) CreateCustomer(customer mdl.Customer) (mdl.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := customer.Validate()
	if err != nil {
		return mdl.Customer{}, err
	}
	customer.ID = s.nextID
	s.customers = append(s.customers, customer)
	s.nextID++
	return customer, nil
}

func (s *InMemoryCustomerStore) GetCustomer(id int) (mdl.Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, c := range s.customers {
		if c.ID == id {
			return c, nil
		}
	}

	return mdl.Customer{}, mdl.ErrCustomerNotFound

}

func (s *InMemoryCustomerStore) UpdateCustomer(id int, customer mdl.Customer) (mdl.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := customer.Validate()
	if err != nil {
		return mdl.Customer{}, err
	}
	for i, c := range s.customers {
		if c.ID == id {
			customer.ID = id
			s.customers[i] = customer
			return customer, nil
		}
	}

	return mdl.Customer{}, mdl.ErrCustomerNotFound
}

func (s *InMemoryCustomerStore) DeleteCustomer(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.customers {
		if c.ID == id {
			s.customers = append(s.customers[:i], s.customers[i+1:]...)
			return nil
		}
	}

	return mdl.ErrCustomerNotFound
}
