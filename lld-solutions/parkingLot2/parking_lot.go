package main

import "fmt"

// ParkingLot represents the entire parking system
type ParkingLot struct {
	Entrances  []Entrance
	Exits      []Exit
	Display    Display
	Floors     []*ParkingFloor
	Strategies map[string]ParkingStrategy
	Payment    map[string]PaymentProcessor
}

// Entrance represents an entry point in the parking lot
type Entrance struct {
	ID   int
	Name string
}

// Exit represents an exit point in the parking lot
type Exit struct {
	ID   int
	Name string
}

// ParkingFloor represents a single floor in the parking lot
type ParkingFloor struct {
	FloorNumber int
	Slots       map[int]*Slot
}

// Slot represents an individual parking space
type Slot struct {
	SlotID      int
	IsFree      bool
	FloorNumber int
	Type        SlotType
	Vehicle     *Vehicle
}

// SlotType defines the type of parking slot
type SlotType interface {
	GetCost() int
	GetVehicleType() string
}

type MiniSlot struct{}

func (m MiniSlot) GetCost() int {
	return 10
}

func (m MiniSlot) GetVehicleType() string {
	return "Mini"
}

type CompactSlot struct{}

func (c CompactSlot) GetCost() int {
	return 20
}

func (c CompactSlot) GetVehicleType() string {
	return "Compact"
}

type LargeSlot struct{}

func (l LargeSlot) GetCost() int {
	return 30
}

func (l LargeSlot) GetVehicleType() string {
	return "Large"
}

// Vehicle represents a parked vehicle
type Vehicle struct {
	LicensePlate string
	Type         string
}

type Display struct {
	FreeMiniSlots    int
	FreeCompactSlots int
	FreeLargeSlots   int
}

func (d *Display) UpdateDisplay(floors []*ParkingFloor) {
	mini, compact, large := 0, 0, 0
	for _, floor := range floors {
		for _, slot := range floor.Slots {
			if slot.IsFree {
				switch slot.Type.GetVehicleType() {
				case "Mini":
					mini++
				case "Compact":
					compact++
				case "Large":
					large++
				}
			}
		}
	}
	d.FreeMiniSlots = mini
	d.FreeCompactSlots = compact
	d.FreeLargeSlots = large
	fmt.Printf("Mini: %d, Compact: %d, Large: %d\n", mini, compact, large)
}

type ParkingTicket struct {
	TicketID    int
	Vehicle     Vehicle
	Slot        *Slot
	EntryTime   int64
	ExitTime    int64
	TotalAmount float64
	Paid        bool
}

type ParkingStrategy interface {
	FindSlot(floors []*ParkingFloor, vehicleType string) *Slot
}

type NearestFirst struct{}

func (n NearestFirst) FindSlot(floors []*ParkingFloor, vehicleType string) *Slot {
	for _, floor := range floors {
		for _, slot := range floor.Slots {
			if slot.IsFree && slot.Type.GetVehicleType() == vehicleType {
				return slot
			}
		}
	}
	return nil
}

type FarthestFirst struct{}

func (f FarthestFirst) FindSlot(floors []*ParkingFloor, vehicleType string) *Slot {
	for i := len(floors) - 1; i >= 0; i-- {
		for _, slot := range floors[i].Slots {
			if slot.IsFree && slot.Type.GetVehicleType() == vehicleType {
				return slot
			}
		}
	}
	return nil
}

type PaymentProcessor interface {
	Charge(ticket *ParkingTicket) bool
}

type CashPayment struct{}

func (c CashPayment) Charge(ticket *ParkingTicket) bool {
	fmt.Println("Payment received in cash:", ticket.TotalAmount)
	ticket.Paid = true
	return true
}

type CreditCardPayment struct{}

func (cc CreditCardPayment) Charge(ticket *ParkingTicket) bool {
	fmt.Println("Payment received via credit card:", ticket.TotalAmount)
	ticket.Paid = true
	return true
}

func (p *ParkingLot) AddEntrance(name string) {
	id := len(p.Entrances) + 1
	p.Entrances = append(p.Entrances, Entrance{ID: id, Name: name})
}

func (p *ParkingLot) AddExit(name string) {
	id := len(p.Exits) + 1
	p.Exits = append(p.Exits, Exit{ID: id, Name: name})
}

func (p *ParkingLot) ParkVehicle(vehicle Vehicle, strategy ParkingStrategy) *ParkingTicket {
	slot := strategy.FindSlot(p.Floors, vehicle.Type)
	if slot == nil {
		fmt.Println("No parking slot available")
		return nil
	}

	slot.IsFree = false
	slot.Vehicle = &vehicle

	ticket := &ParkingTicket{
		TicketID:  len(p.Floors) + 1,
		Vehicle:   vehicle,
		Slot:      slot,
		EntryTime: 0, // Assume timestamp
		Paid:      false,
	}
	fmt.Printf("Vehicle %s parked at Floor %d, Slot %d\n", vehicle.LicensePlate, slot.FloorNumber, slot.SlotID)
	return ticket
}
func (p *ParkingLot) ExitVehicle(ticket *ParkingTicket, payment PaymentProcessor) {
	if ticket == nil || ticket.Paid {
		fmt.Println("Invalid ticket or already paid")
		return
	}

	// Calculate charges (assuming 1-hour parking for simplicity)
	ticket.TotalAmount = float64(ticket.Slot.Type.GetCost())
	payment.Charge(ticket)

	// Free up the slot
	ticket.Slot.IsFree = true
	ticket.Slot.Vehicle = nil
}
