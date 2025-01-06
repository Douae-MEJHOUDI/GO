package main

import (
	"02-area-calculator/areas"
	"fmt"
)

func main() {
	var x, y float64 = 5, 6
	var area float64 = areas.Area(x, y)
	fmt.Println(area)
	fmt.Println(int(areas.Meter_to_foot(area)))

}
