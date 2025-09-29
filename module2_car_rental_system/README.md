# Car rental system
# Structs
I used structs to define the key entities:
* Customer - represents a customer with name, Id, Contact, driver license
* Car - holds car details (make, model, year, license, plate number, rental price per day)
* Reservation - hold reservation details (ID, carId, customerId, start date, end date)
# Interfaces
I created interfaces to define reusable behaviors:
* CarService - defines available cars, get all cars, add a new car
* ReservationService - defines to add new reservation, modify existing one, cancel reservation, is available based on start date and end date, get all reservations

#  Methods
Each struct has associated methods. For example:
* 	
    1. Add(Id int, carId int, customer_Id int, start_date string, end_date string) error
	2. modifyReservation(reservationId int, start_date string, end_date string)
	3. cancelReservation(Id int)
	4. IsAvailable(carId int, b_start, b_end string) bool
	5. GetAll() []Reservation
* 
    1. GetAllCars() []Car
	2. Add(ID int, make string, model string, year int, license string, plate_number string,  price int)
	4. AvailableCars(carType string, maxPrice int) []Car


#  How to Run Locally
1. Clone the repository:
git clone https://github.com/AkhilaGuda/GoActivities.git
cd GoActivities/module2/module2_car_rental_system