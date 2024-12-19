package vehicles

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	GetType() string
}

type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

type Motorcycle struct {
	Speed    float64
	Type     string
	FuelType string
}

func (c Car) GetType() string {
	return "main.Car"
}

func (m Motorcycle) GetType() string {
	return "main.Motorcycle"
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	result := make(map[string]float64)
	for _, v := range vehicles {
		result[v.GetType()] = v.CalculateTravelTime(distance)
	}
	return result
}
