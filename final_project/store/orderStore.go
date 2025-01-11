package store

import (
	mdl "final_project/models"
)

type OrderStore interface {
	CreateOrder(order mdl.Order) (mdl.Order, error)
	GetOrder(id int) (mdl.Order, error)
	UpdateOrder(id int, order mdl.Order) (mdl.Order, error)
	DeleteOrder(id int) error
	GetAllOrders() ([]mdl.Order, error)
}
