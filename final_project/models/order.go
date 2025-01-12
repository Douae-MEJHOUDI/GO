package models

import (
	"errors"
	"time"
)

var (
	ErrOrderNotFound         = errors.New("Order not found")
	ErrEmptyOrder            = errors.New("empty order")
	ErrInvalidQuantity       = errors.New("invalid quantity")
	ErrNoOrdersFound         = errors.New("No order exist yet")
	ErrOrderNotSavedInMemory = errors.New("order changes were not saved into memory")
)

type Order struct {
	ID         int         `json:"id"`
	Customer   Customer    `json:"customer"`
	Items      []OrderItem `json:"items"`
	TotatPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `json:"created_at"`
	Status     string      `json:"status"`
}

type OrderItem struct {
	Book     Book `json:"book"`
	Quantity int  `json:"quantity"`
}

func (o *Order) Validate() error {
	err := o.Customer.Validate()
	if err != nil {
		return err
	}
	if len(o.Items) == 0 {
		return ErrEmptyOrder
	}
	for _, item := range o.Items {
		err := item.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

func (oi *OrderItem) Validate() error {
	err := oi.Book.Validate()
	if err != nil {
		return err
	}
	if oi.Quantity <= 0 {
		return ErrInvalidQuantity
	}
	return nil
}
