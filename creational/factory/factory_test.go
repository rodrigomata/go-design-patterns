package factory

import (
	"testing"
)

func TestVehicleCar(t *testing.T) {
	car, err := GetVehicle(Car)
	if err != nil {
		t.Fatal("A Vehicle of type 'Car' must exist")
	}

	canDrive := car.Drive(4)
	if canDrive != true {
		t.Error("A car only allows up to 4 passengers")
	}
}

func TestVehicleBike(t *testing.T) {
	bike, err := GetVehicle(Bike)
	if err != nil {
		t.Fatal("A Vehicle of type 'Bike' must exist")
	}

	canDrive := bike.Drive(2)
	if canDrive != true {
		t.Error("A bike only allows up to 2 passengers")
	}
}

func TestVehicleNonExistent(t *testing.T) {
	_, err := GetVehicle(50)
	if err == nil {
		t.Fatal("A Vehicle with ID 50 must return an error")
	}
	t.Log("LOG:", err)
}
