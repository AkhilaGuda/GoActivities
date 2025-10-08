package db

import (
	"fmt"
	"time"
)

func (db DBManager) GetAll() []Reservation {
	return db.reservations
}
func (db *DBManager) AddReservation(Id int, carId int, customer_Id int, start_date string, end_date string) error {
	newReservation := Reservation{
		Id:          Id,
		CarId:       carId,
		Customer_Id: customer_Id,
		Start_date:  start_date,
		End_date:    end_date,
	}
	layout := "2006-01-02"
	new_startDate, _ := time.Parse(layout, start_date)
	new_endDate, _ := time.Parse(layout, end_date)

	for _, existing := range db.reservations {
		if existing.CarId == carId {
			existing_startDate, _ := time.Parse(layout, existing.Start_date)
			existing_endDate, _ := time.Parse(layout, existing.End_date)
			if new_startDate.Before(existing_endDate) && new_endDate.After(existing_startDate) {
				return fmt.Errorf("car %d is already reserved for these dates", carId)
			}
		}
	}
	(*db).reservations = append((*db).reservations, newReservation)
	return nil

}

func (db *DBManager) ModifyReservation(reservationId int, start_date string, end_date string) {
	for i, exisiting := range (*db).reservations {
		if exisiting.Id == reservationId {
			(*db).reservations[i].Start_date = start_date
			(*db).reservations[i].End_date = end_date
			break
		}
	}

}

func (db *DBManager) CancelReservation(Id int) {
	for i, res := range (*db).reservations {
		if res.Id == Id {
			(*db).reservations = append((*db).reservations[:i], (*db).reservations[i+1:]...)
		}
	}
}
