package db

import "time"

func (db DBManager) GetAllCars() []Car {
	var result []Car
	for _, car := range db.cars {
		result = append(result, car)
	}
	return result
}
func (db *DBManager) AddCar(ID int, make string, model string, year int, license string, plate_number string, price int) {
	db.cars = append(db.cars, Car{Make: make, Model: model, Year: year, License: license, Plate_number: plate_number, Rental_price_per_day: price})
}

func (db DBManager) AvailableCars(carType string, maxPrice int) []Car {
	var result []Car
	for _, car := range db.cars {
		if car.Make == carType && car.Rental_price_per_day <= maxPrice {
			result = append(result, car)
		}
	}
	return result
}
func (db DBManager) IsAvailable(carId int, b_start, b_end string) bool {
	layout := "2006-01-02"
	b_startDate, _ := time.Parse(layout, b_start)
	b_endDate, _ := time.Parse(layout, b_end)
	for _, res := range db.reservations {
		if res.CarId == carId {
			a_startDate, _ := time.Parse(layout, res.Start_date)
			a_endDate, _ := time.Parse(layout, res.End_date)

			//a existing reservation, b new booking request
			if b_startDate.Before(a_endDate) && b_endDate.After(a_startDate) {
				return false
			}
		}

	}
	return true
}
