package store

import (
	mdl "final_project/models"
	"sync"
)

type InMemoryOrderStore struct {
	mu     sync.RWMutex
	orders []mdl.Order
	nextID int
}

func NewOrderStore() *InMemoryOrderStore {
	return &InMemoryOrderStore{
		orders: []mdl.Order{},
		nextID: 1,
	}
}

func (s *InMemoryOrderStore) CreateOrder(order mdl.Order) (mdl.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := order.Validate()
	if err != nil {
		return mdl.Order{}, err
	}
	order.ID = s.nextID
	s.orders = append(s.orders, order)
	s.nextID++
	return order, nil
}

func (s *InMemoryOrderStore) GetOrder(id int) (mdl.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, c := range s.orders {
		if c.ID == id {
			return c, nil
		}
	}

	return mdl.Order{}, mdl.ErrOrderNotFound

}

func (s *InMemoryOrderStore) UpdateOrder(id int, order mdl.Order) (mdl.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := order.Validate()
	if err != nil {
		return mdl.Order{}, err
	}
	for i, c := range s.orders {
		if c.ID == id {
			order.ID = id
			s.orders[i] = order
			return order, nil
		}
	}

	return mdl.Order{}, mdl.ErrOrderNotFound
}

func (s *InMemoryOrderStore) DeleteOrder(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.orders {
		if c.ID == id {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			return nil
		}
	}

	return mdl.ErrOrderNotFound
}

func (s *InMemoryOrderStore) GetAllOrders() ([]mdl.Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	//orders := s.orders

	return s.orders, nil

}
