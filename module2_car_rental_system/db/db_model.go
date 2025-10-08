package db

type Car struct {
	Make                 string
	Model                string
	Year                 int
	License              string
	Plate_number         string
	Rental_price_per_day int
}
type Reservation struct {
	Id          int
	CarId       int
	Customer_Id int
	Start_date  string
	End_date    string
}
type Customer struct {
	CustomerId    int
	CustomerName  string
	Contact       string
	DriverLicense string
}
type DBManager struct {
	cars         []Car
	reservations []Reservation
}
type Service interface {
	//car methods
	GetAllCars() []Car
	AddCar(ID int, make string, model string, year int, license string, plate_number string, price int)
	AvailableCars(carType string, maxPrice int) []Car
	//reservation methods
	AddReservation(Id int, carId int, customer_Id int, start_date string, end_date string) error
	ModifyReservation(reservationId int, start_date string, end_date string)
	CancelReservation(Id int)
	IsAvailable(carId int, b_start, b_end string) bool
	GetAll() []Reservation
}
