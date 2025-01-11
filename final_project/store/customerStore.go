package store

import (
	mdl "final_project/models"
)

type CustomerStore interface {
	CreateCustomer(customer mdl.Customer) (mdl.Customer, error)
	GetCustomer(id int) (mdl.Customer, error)
	UpdateCustomer(id int, customer mdl.Customer) (mdl.Customer, error)
	DeleteCustomer(id int) error
}
