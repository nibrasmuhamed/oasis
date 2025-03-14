package main

func main() {
	// Create parking lot
	parkingLot := &ParkingLot{
		Floors: []*ParkingFloor{
			{FloorNumber: 1, Slots: map[int]*Slot{
				1: {SlotID: 1, IsFree: true, FloorNumber: 1, Type: MiniSlot{}},
				2: {SlotID: 2, IsFree: true, FloorNumber: 1, Type: CompactSlot{}},
			}},
		},
		Strategies: map[string]ParkingStrategy{
			"Nearest":  NearestFirst{},
			"Farthest": FarthestFirst{},
		},
		Payment: CashPayment{},
	}

	// Add entrance and exit
	parkingLot.AddEntrance("Main Gate")
	parkingLot.AddExit("Exit 1")

	// Park a vehicle
	vehicle := Vehicle{LicensePlate: "KA-01-1234", Type: "Mini"}
	ticket := parkingLot.ParkVehicle(vehicle, parkingLot.Strategies["Nearest"])

	// Process exit and payment
	if ticket != nil {
		parkingLot.ExitVehicle(ticket, parkingLot.Payment)
	}
}
