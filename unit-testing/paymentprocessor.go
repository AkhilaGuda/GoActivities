package main

type PaymentProcessor interface {
	Charge(amount float64) error
}

type Order struct {
	Processor PaymentProcessor
	Total     float64
}

func (o *Order) Checkout() error {
	return o.Processor.Charge(o.Total)
}
