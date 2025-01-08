package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Vehicule struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicule
	NumberOfDoors int
}

type Insurable interface {
	CalculateInsurance() int
}

type Printable interface {
	Details()
}

func (c Car) Details() {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Number of Doors: %d", c.Make, c.Model, c.Year, c.NumberOfDoors)
}

func (c Car) CalculateInsurance() int {
	return 1000
}

type Truck struct {
	Vehicule
	PayloadCapacity int
}

func (t Truck) Details() {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Payload Capacity: %d", t.Make, t.Model, t.Year, t.PayloadCapacity)
}

func (t Truck) CalculateInsurance() int {
	return 2000
}

func main() {
	car := Car{
		Vehicule: Vehicule{
			Make:  "A",
			Model: "A",
			Year:  2015,
		},
		NumberOfDoors: 4,
	}

	truck := Truck{
		Vehicule: Vehicule{
			Make:  "B",
			Model: "B",
			Year:  2015,
		},
		PayloadCapacity: 2000,
	}

	truck2 := Truck{
		Vehicule: Vehicule{
			Model: "C",
			Make:  "C",
			Year:  2015,
		},
		PayloadCapacity: 3000,
	}

	car.Details()
	fmt.Println()
	truck.Details()
	fmt.Println()

	PrintAll([]Printable{car, truck, truck2})

	content := []Insurable{car, truck, truck2}
	data, _ := json.MarshalIndent(content, " ", "  ")
	os.WriteFile("data.json", data, 0644)

}

func PrintAll(p []Printable) {
	for i, v := range p {
		fmt.Println("Details of vehicule N", i+1)
		v.Details()
		fmt.Println()
	}
}
