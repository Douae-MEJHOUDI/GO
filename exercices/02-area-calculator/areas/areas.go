package areas

const foot float64 = 3.28

func Area(x float64, y float64) float64 {
	return x * y
}

func Meter_to_foot(x float64) float64 {

	return x * foot
}
