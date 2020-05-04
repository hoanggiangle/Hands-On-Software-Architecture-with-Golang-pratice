package main

import (
	"fmt"
)

//  Reservation and Invoice as two generic products
//type Reservation interface{}
type Invoice interface{}

// AbstractFactory is the abstract factory which will create two products - both which need to work with each other
type AbstractFactory interface {
	CreateReservation() int
	CreateInvoice() Invoice
}

// HotelFactory implements AbstractFactory and creates products for the Hotel family of products
type HotelFactory struct{}

func (f HotelFactory) CreateReservation() int {
	return 2
}

func (f HotelFactory) CreateInvoice() Invoice {
	return new(HotelInvoice)
}

// FlightFactory implements AbstractFactory and creates products for the Flight family of products
type FlightFactory struct{}

func (f FlightFactory) CreateReservation() int {
	return 1
}

func (f FlightFactory) CreateInvoice() Invoice {
	return new(FlightReservation)
}

// GetFactory returns a factory given a config
func GetFactory(vertical string) AbstractFactory {
	var factory AbstractFactory
	switch vertical {
	case "flight":
		factory = FlightFactory{}
	case "hotel1":
		factory = HotelFactory{}
	}

	return factory
}

// The code below demonstrates usage of the factories
//

type HotelReservation struct{}
type HotelInvoice struct{}
type FlightReservation struct{}
type FlightInvoice struct{}

func main() {
	hotelFactory := GetFactory("hotel1")
	reservation := hotelFactory.CreateReservation()
	invoice := hotelFactory.CreateInvoice()

	fmt.Printf("%T\n", reservation)
	fmt.Printf("%T\n", invoice)

}
