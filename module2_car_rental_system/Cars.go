package main

type Car struct {
	make                 string
	model                string
	year                 int
	license              string
	plate_number         string
	rental_price_per_day int
}
type CarService interface {
	GetAllCars() []Car
	Add(ID int, make string, model string, year int, license string, plate_number string, price int)
	AvailableCars(carType string, maxPrice int) []Car
}

type Cars map[int]Car

func (c Cars) GetAllCars() []Car {
	var result []Car
	for _, car := range c {
		result = append(result, car)
	}
	return result
}
func (c Cars) Add(ID int, make string, model string, year int, license string, plate_number string, price int) {
	c[ID] = Car{make: make, model: model, year: year, license: license, plate_number: plate_number, rental_price_per_day: price}
}

func (c Cars) AvailableCars(carType string, maxPrice int) []Car {
	var result []Car
	for _, car := range c {
		if car.make == carType && car.rental_price_per_day <= maxPrice {
			result = append(result, car)
		}
	}
	return result
}
