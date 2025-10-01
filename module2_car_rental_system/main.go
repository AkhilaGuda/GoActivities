package main

import (
	"fmt"
)

func main() {
	var carService CarService = Cars{}
	carService.Add(1, "Ford", "Ecosport", 2021, "LMNV", "TS071234", 2000)
	carService.Add(4, "Ford", "Ecosport2", 2022, "LMNVp", "TS071334", 3000)
	carService.Add(2, "Toyota", "Corolla", 2020, "ABC1234", "KA051234", 1500)
	carService.Add(3, "Honda", "Civic", 2019, "XYZ5678", "AP125678", 1800)
	fmt.Println("All car details -")
	for _, car := range carService.GetAllCars() {
		fmt.Println(car)
	}
	fmt.Println()
	var reservationService ReservationService = &Reservations{}
	reservationService.Add(1, 1, 1, "2024-01-01", "2024-01-10")
	reservationService.Add(1, 2, 2, "2024-01-19", "2024-02-20")
	reservationService.Add(1, 3, 3, "2024-01-01", "2024-01-10")
	reservationService.Add(10, 3, 3, "2024-03-09", "2024-03-11")

	availableCars := carService.AvailableCars("Ford", 3000)
	fmt.Println("Available cars for : type-Ford price-3000 :\n", availableCars)

	fmt.Println("\nAvailable cars for specific dates: 2024-01-19 & 2024-02-20 ")
	for id, vehicle := range carService.GetAllCars() {
		if reservationService.IsAvailable(id, "2024-01-19", "2024-02-20") {
			fmt.Println(vehicle.model)
		}
	}

	fmt.Println("\nChecking double reservation for same car")
	fmt.Println(reservationService.Add(1, 3, 3, "2024-01-01", "2024-01-10"))
	if reservationService.Add(1, 3, 3, "2024-01-20", "2024-01-30") == nil {
		fmt.Println("Car Reserved successfully")
	}

	fmt.Println("Updated Reservations after duplicate entry: ")
	for _, reservation := range reservationService.GetAll() {
		fmt.Println(reservation)
	}

	reservationService.modifyReservation(10, "2024-01-02", "2024-01-10")
	fmt.Println("\n After modification :")
	for _, reservation := range reservationService.GetAll() {
		fmt.Println(reservation)
	}
	fmt.Println("\n Deleting reservation Id : 10")
	reservationService.cancelReservation(10)
	for _, reservation := range reservationService.GetAll() {
		fmt.Println(reservation)
	}
}
