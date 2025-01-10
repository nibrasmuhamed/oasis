package main

import (
	"fmt"
)

// VehicleType represents types of vehicles
type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Bus
)

// Vehicle represents a vehicle with its properties
type Vehicle struct {
	LicensePlate string
	Type         VehicleType
}

// ParkingSpot represents a single parking spot
type ParkingSpot struct {
	SpotNumber int
	Level      int
	Vehicle    *Vehicle
}

// ParkingLot represents the entire parking lot with multiple levels and spots
type ParkingLot struct {
	Levels        int
	SpotsPerLevel int
	Spots         [][]*ParkingSpot
}

// NewParkingLot creates a new parking lot with specified levels and spots per level
func NewParkingLot(levels, spotsPerLevel int) *ParkingLot {
	parkingLot := &ParkingLot{
		Levels:        levels,
		SpotsPerLevel: spotsPerLevel,
		Spots:         make([][]*ParkingSpot, levels),
	}
	for i := range parkingLot.Spots {
		parkingLot.Spots[i] = make([]*ParkingSpot, spotsPerLevel)
	}
	return parkingLot
}

// ParkVehicle parks a vehicle in the parking lot, returns true if successful
func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) bool {
	for level := 0; level < pl.Levels; level++ {
		for spot := 0; spot < pl.SpotsPerLevel; spot++ {
			if pl.Spots[level][spot] == nil {
				pl.Spots[level][spot] = &ParkingSpot{
					SpotNumber: spot + 1,
					Level:      level + 1,
					Vehicle:    vehicle,
				}
				return true
			}
		}
	}
	return false
}

// UnparkVehicle removes a vehicle from the parking lot, returns true if successful
func (pl *ParkingLot) UnparkVehicle(vehicle *Vehicle) bool {
	for level := 0; level < pl.Levels; level++ {
		for spot := 0; spot < pl.SpotsPerLevel; spot++ {
			if pl.Spots[level][spot] != nil && pl.Spots[level][spot].Vehicle == vehicle {
				pl.Spots[level][spot] = nil
				return true
			}
		}
	}
	return false
}

func main() {
	pl := NewParkingLot(2, 5)
	fmt.Println("Parking lot created with 2 levels and 5 spots per level.")

	vehicle1 := &Vehicle{LicensePlate: "ABC123", Type: Car}
	vehicle2 := &Vehicle{LicensePlate: "XYZ789", Type: Motorcycle}

	if pl.ParkVehicle(vehicle1) {
		fmt.Printf("Vehicle %s parked successfully.\n", vehicle1.LicensePlate)
	} else {
		fmt.Printf("Parking failed for vehicle %s.\n", vehicle1.LicensePlate)
	}

	if pl.ParkVehicle(vehicle2) {
		fmt.Printf("Vehicle %s parked successfully.\n", vehicle2.LicensePlate)
	} else {
		fmt.Printf("Parking failed for vehicle %s.\n", vehicle2.LicensePlate)
	}

	if pl.UnparkVehicle(vehicle1) {
		fmt.Printf("Vehicle %s unparked successfully.\n", vehicle1.LicensePlate)
	} else {
		fmt.Printf("Unparking failed for vehicle %s.\n", vehicle1.LicensePlate)
	}
}
