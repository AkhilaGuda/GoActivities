package car

type Car struct {
	Make                 string
	Model                string
	Year                 int
	License              string
	Plate_number         string
	Rental_price_per_day int
}
type CarService interface {
	GetAllCars() []Car
	Add(ID int, make string, model string, year int, license string, plate_number string, price int)
	AvailableCars(carType string, maxPrice int) []Car
}

type Cars map[int]Car

func NewCars() CarService {
	return Cars{}
}
func (c Cars) GetAllCars() []Car {
	var result []Car
	for _, car := range c {
		result = append(result, car)
	}
	return result
}
func (c Cars) Add(ID int, make string, model string, year int, license string, plate_number string, price int) {
	c[ID] = Car{Make: make, Model: model, Year: year, License: license, Plate_number: plate_number, Rental_price_per_day: price}
}

func (c Cars) AvailableCars(carType string, maxPrice int) []Car {
	var result []Car
	for _, car := range c {
		if car.Make == carType && car.Rental_price_per_day <= maxPrice {
			result = append(result, car)
		}
	}
	return result
}
