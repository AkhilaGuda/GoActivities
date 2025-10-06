package reservation

import (
	"fmt"
	"time"
)

type Reservation struct {
	Id          int
	CarId       int
	Customer_Id int
	Start_date  string
	End_date    string
}

type ReservationService interface {
	Add(Id int, carId int, customer_Id int, start_date string, end_date string) error
	ModifyReservation(reservationId int, start_date string, end_date string)
	CancelReservation(Id int)
	IsAvailable(carId int, b_start, b_end string) bool
	GetAll() []Reservation
}

type Reservations []Reservation

func (r Reservations) GetAll() []Reservation {
	return r
}
func (r *Reservations) Add(Id int, carId int, customer_Id int, start_date string, end_date string) error {
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

	for _, existing := range *r {
		if existing.CarId == carId {
			existing_startDate, _ := time.Parse(layout, existing.Start_date)
			existing_endDate, _ := time.Parse(layout, existing.End_date)
			if new_startDate.Before(existing_endDate) && new_endDate.After(existing_startDate) {
				return fmt.Errorf("car %d is already reserved for these dates", carId)
			}
		}
	}
	*r = append(*r, newReservation)
	return nil

}

func (r *Reservations) ModifyReservation(reservationId int, start_date string, end_date string) {
	for i, exisiting := range *r {
		if exisiting.Id == reservationId {
			(*r)[i].Start_date = start_date
			(*r)[i].End_date = end_date
			break
		}
	}

}

func (r *Reservations) CancelReservation(Id int) {
	for i, res := range *r {
		if res.Id == Id {
			*r = append((*r)[:i], (*r)[i+1:]...)
		}
	}
}

func (r Reservations) IsAvailable(carId int, b_start, b_end string) bool {
	layout := "2006-01-02"
	b_startDate, _ := time.Parse(layout, b_start)
	b_endDate, _ := time.Parse(layout, b_end)
	for _, res := range r {
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
