package main

import (
	"fmt"
	"time"
)

type Reservation struct {
	Id          int
	carId       int
	customer_Id int
	start_date  string
	end_date    string
}

type ReservationService interface {
	Add(Id int, carId int, customer_Id int, start_date string, end_date string) error
	modifyReservation(reservationId int, start_date string, end_date string)
	cancelReservation(Id int)
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
		carId:       carId,
		customer_Id: customer_Id,
		start_date:  start_date,
		end_date:    end_date,
	}
	layout := "2006-01-02"
	new_startDate, _ := time.Parse(layout, start_date)
	new_endDate, _ := time.Parse(layout, end_date)

	for _, existing := range *r {
		if existing.carId == carId {
			existing_startDate, _ := time.Parse(layout, existing.start_date)
			existing_endDate, _ := time.Parse(layout, existing.end_date)
			if new_startDate.Before(existing_endDate) && new_endDate.After(existing_startDate) {
				return fmt.Errorf("car %d is already reserved for these dates", carId)
			}
		}
	}
	*r = append(*r, newReservation)
	return nil

}

func (r *Reservations) modifyReservation(reservationId int, start_date string, end_date string) {
	// res.start_date = start_date
	// res.end_date = end_date
	for i, exisiting := range *r {
		if exisiting.Id == reservationId {
			(*r)[i].start_date = start_date
			(*r)[i].end_date = end_date
			break
		}
	}

}

func (r *Reservations) cancelReservation(Id int) {
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
		if res.carId == carId {
			a_startDate, _ := time.Parse(layout, res.start_date)
			a_endDate, _ := time.Parse(layout, res.end_date)

			//a existing reservation, b new booking request
			if b_startDate.Before(a_endDate) && b_endDate.After(a_startDate) {
				return false
			}
		}

	}
	return true
}
