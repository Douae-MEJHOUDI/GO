package store

import (
	"errors"
	mdl "final_project/models"
	"fmt"
	"log"
	"sync"
	"time"
)

type InMemoryOrderStore struct {
	mu        sync.RWMutex
	orders    []mdl.Order
	books     BookStore
	customers CustomerStore
	nextID    int
	data      *DataManager
}

type DataMOrder struct {
	Orders []mdl.Order `json:"orders"`
	NextID int         `json:"next_id"`
}

func NewOrderStore(books BookStore, customerstore CustomerStore, data *DataManager) *InMemoryOrderStore {
	store := &InMemoryOrderStore{
		orders:    []mdl.Order{},
		books:     books,
		customers: customerstore,
		nextID:    1,
		data:      data,
	}

	var dataM DataMOrder

	err := data.LoadData("orders.json", &dataM)
	if err != nil {
		log.Println("Couldn't load order data: ", err)
	} else if len(dataM.Orders) > 0 {
		store.orders = dataM.Orders
		store.nextID = dataM.NextID
	}

	return store
}

func (s *InMemoryOrderStore) saveOrderData() error {
	var dataM DataMOrder
	dataM.Orders = s.orders
	dataM.NextID = s.nextID
	return s.data.SaveData("orders.json", dataM)
}

func (s *InMemoryOrderStore) CreateOrder(order mdl.Order) (mdl.Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	err := order.Validate()
	if err != nil {
		return mdl.Order{}, err
	}
	cust, err := s.customers.GetCustomer(order.Customer.ID)
	if err != nil {
		return mdl.Order{}, errors.New("invalid customer: " + err.Error())
	}

	order.Customer = cust
	var totalPrice float64

	for i, item := range order.Items {
		book, err := s.books.GetBook(item.Book.ID)
		if err != nil {
			return mdl.Order{}, errors.New("invalid book item: " + err.Error())
		}

		if book.Stock < item.Quantity {
			return mdl.Order{}, fmt.Errorf("insufficient stock for book %s: exist %d, wanted %d", book.Title, book.Stock, item.Quantity)
		}
		order.Items[i].Book = book
		totalPrice += float64(item.Quantity) * book.Price
	}

	for _, item := range order.Items {
		book := item.Book
		book.Stock -= item.Quantity
		_, err := s.books.UpdateBook(book.ID, book)
		if err != nil {
			return mdl.Order{}, errors.New("failef to update book stock: " + err.Error())
		}
	}

	order.ID = s.nextID
	order.TotatPrice = totalPrice
	order.CreatedAt = time.Now()
	order.Status = "successful"

	s.orders = append(s.orders, order)
	s.nextID++

	err = s.saveOrderData()
	if err != nil {
		order.Status = "Unsaved"
		return mdl.Order{}, mdl.ErrOrderNotSavedInMemory
	}

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
			err = s.saveOrderData()
			if err != nil {
				order.Status = "Unsaved"
				return mdl.Order{}, mdl.ErrOrderNotSavedInMemory
			}
			return order, nil
		}
	}

	return mdl.Order{}, mdl.ErrOrderNotFound
}

func (s *InMemoryOrderStore) DeleteOrder(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, o := range s.orders {
		if o.ID == id {
			s.orders = append(s.orders[:i], s.orders[i+1:]...)
			err := s.saveOrderData()
			if err != nil {
				o.Status = "Unsaved"
				return mdl.ErrOrderNotSavedInMemory
			}
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
