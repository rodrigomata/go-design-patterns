package factory

import "fmt"

// Vehicle type exposed
type Vehicle interface {
	Drive(passengers int) bool
}

// Types of Vehicles
const (
	Car  = 1
	Bike = 2
)

// GetVehicle checks if there is an implementation of a Vehicle
func GetVehicle(v int) (Vehicle, error) {
	switch v {
	case Car:
		return &CarF{}, nil
	case Bike:
		return &BikeF{}, nil
	default:
		return nil, fmt.Errorf("Vehicle %d not supported", v)
	}
}

// CarF is the struct of a Car
type CarF struct{}

// BikeF is the struct of a Bike
type BikeF struct{}

// Drive lets you know if the total passengers can drive in a Car
func (c *CarF) Drive(passengers int) bool {
	return passengers <= 4
}

// Drive lets you know if the total passengers can drive in a Bike
func (c *BikeF) Drive(passengers int) bool {
	return passengers <= 2
}
