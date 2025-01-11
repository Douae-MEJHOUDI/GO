package service

import (
	"encoding/json"
	mdl "final_project/models"
	store "final_project/store"
	"fmt"
	"os"
	"time"
)

type Generator struct {
	stores *store.Stores
}

func NewGenerator(stores *store.Stores) *Generator {
	return &Generator{
		stores: stores,
	}
}

func (g *Generator) GenerateReport() error {
	endPeriod := time.Now()
	startPeriod := endPeriod.Add(-24 * time.Hour)

	orders, err := g.stores.Orders.GetAllOrders()
	if err != nil {
		return fmt.Errorf("failed to get orders: %v", err)
	}

	var totalRevenue float64

	bookSales := make(map[int]mdl.BookSales)

	for _, order := range orders {

		if order.CreatedAt.After(startPeriod) && order.CreatedAt.Before(endPeriod) {
			totalRevenue += order.TotatPrice
			for _, item := range order.Items {
				sales := bookSales[item.Book.ID]
				sales.Book = item.Book
				sales.QuantitySold += item.Quantity
				sales.Revenue += float64(item.Quantity) * item.Book.Price
				bookSales[item.Book.ID] = sales
			}
		}
	}

	var topBooks []mdl.BookSales
	for _, sales := range bookSales {
		topBooks = append(topBooks, sales)
	}

	report := mdl.SalesReport{
		ID:           fmt.Sprintf("REPORT_%s", time.Now().Format("20060102150405")),
		StartPeriod:  startPeriod,
		EndPeriod:    endPeriod,
		GeneratedAt:  time.Now(),
		TotalRevenue: totalRevenue,
		TotalOrders:  len(orders),
		TopBooks:     topBooks,
	}

	fileName := fmt.Sprintf("report_%s.json", time.Now().Format("20060102150405"))

	data, err := json.MarshalIndent(report, " ", "  ")

	if err != nil {
		return fmt.Errorf("failed to write report: %v", err)
	}
	err = os.WriteFile("output-reports/"+fileName, data, 0644)

	if err != nil {
		return fmt.Errorf("failed to write report: %v", err)
	}

	return nil
}

func (g *Generator) Start() {
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()
	done := make(chan struct{})
	defer close(done)

	err := g.GenerateReport()
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case <-ticker.C:
			err := g.GenerateReport()
			if err != nil {
				fmt.Println(err)
			}
		case <-done:
			return
		}
	}
}
