package models

import "time"

type BookSales struct {
	Book         Book    `json:"book"`
	QuantitySold int     `json:"quantity_sold"`
	Revenue      float64 `json:"revenue"`
}

type SalesReport struct {
	ID           string      `json:"id"`
	StartPeriod  time.Time   `json:"start_period"`
	EndPeriod    time.Time   `json:"end_period"`
	GeneratedAt  time.Time   `json:"generated_at"`
	TotalRevenue float64     `json:"total_revenue"`
	TotalOrders  int         `json:"total_orders"`
	TopBooks     []BookSales `json:"top_selling_books"`
}
