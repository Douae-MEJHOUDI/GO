package store

import (
	mdl "final_project/models"
	"log"
	"sync"
	"time"
)

type InMemoryCustomerStore struct {
	mu        sync.RWMutex
	customers []mdl.Customer
	data      *DataManager
	nextID    int
	orders    OrderStore
}

type DataMCustomer struct {
	Customers []mdl.Customer `json:"customer"`
	NextID    int            `json:"next_id"`
}

func NewCustomerStore(data *DataManager, orders OrderStore) *InMemoryCustomerStore {
	store := &InMemoryCustomerStore{
		customers: []mdl.Customer{},
		orders:    orders,
		nextID:    1,
		data:      data,
	}

	var dataM DataMCustomer
	err := data.LoadData("customers.json", &dataM)
	if err != nil {
		log.Println("failed to load customer data: ", err.Error())
	} else if len(dataM.Customers) > 0 {
		store.customers = dataM.Customers
		store.nextID = dataM.NextID
	}

	return store
}

func (s *InMemoryCustomerStore) saveCustomerData() error {
	var dataM DataMCustomer
	dataM.Customers = s.customers
	dataM.NextID = s.nextID

	return s.data.SaveData("customers.json", dataM)
}

func (s *InMemoryCustomerStore) CreateCustomer(customer mdl.Customer) (mdl.Customer, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := customer.Validate()
	if err != nil {
		return mdl.Customer{}, err
	}
	customer.ID = s.nextID
	customer.CreatedAt = time.Now()
	s.customers = append(s.customers, customer)
	s.nextID++

	err = s.saveCustomerData()
	if err != nil {
		return mdl.Customer{}, mdl.ErrCustomerNotSavedInMemory
	}
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
			customer.CreatedAt = c.CreatedAt
			s.customers[i] = customer
			err = s.saveCustomerData()
			if err != nil {
				s.customers[i] = c
				return mdl.Customer{}, mdl.ErrCustomerNotSavedInMemory
			}

			return customer, nil
		}
	}

	return mdl.Customer{}, mdl.ErrCustomerNotFound
}

func (s *InMemoryCustomerStore) DeleteCustomer(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.orders != nil {
		orders, err := s.orders.GetAllOrders()
		if err != nil {
			return err
		}

		for _, order := range orders {
			if order.Customer.ID == id {
				return mdl.ErrCustomerHasOrders
			}
		}
	}

	for i, c := range s.customers {
		if c.ID == id {
			s.customers = append(s.customers[:i], s.customers[i+1:]...)
			err := s.saveCustomerData()
			if err != nil {
				s.customers[i] = c
				return mdl.ErrCustomerNotSavedInMemory
			}
			return nil
		}
	}

	return mdl.ErrCustomerNotFound
}

func (s *InMemoryCustomerStore) GetAllCustomers() ([]mdl.Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.customers, nil
}
