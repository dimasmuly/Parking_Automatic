package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ParkingSlot struct {
	SlotNumber         int
	RegistrationNumber string
	Color              string 
	IsOccupied         bool
}

type ParkingLot struct {
	Capacity     int
	Slots        []ParkingSlot
	CarSlotMap   map[string]int // Map untuk lookup cepat
}

func NewParkingLot(capacity int) *ParkingLot {
	slots := make([]ParkingSlot, capacity)
	for i := 0; i < capacity; i++ {
		slots[i] = ParkingSlot{
			SlotNumber: i + 1,
			IsOccupied: false,
		}
	}
	return &ParkingLot{
		Capacity:   capacity,
		Slots:      slots,
		CarSlotMap: make(map[string]int), // Initialize map
	}
}

func (pl *ParkingLot) CreateParkingLot(capacity int) {
	pl.Capacity = capacity
	pl.Slots = make([]ParkingSlot, capacity)
	pl.CarSlotMap = make(map[string]int) // Reset map
	for i := 0; i < capacity; i++ {
		pl.Slots[i] = ParkingSlot{
			SlotNumber: i + 1,
			IsOccupied: false,
		}
	}
	fmt.Println("\n=== PARKING LOT CREATED ===")
	fmt.Printf("Created a parking lot with %d slots\n", capacity)
	fmt.Println("============================\n")
}

// Updated Park function dengan warna
func (pl *ParkingLot) Park(registrationNumber, color string) {
	// Find the nearest available slot
	for i := 0; i < pl.Capacity; i++ {
		if !pl.Slots[i].IsOccupied {
			pl.Slots[i].IsOccupied = true
			pl.Slots[i].RegistrationNumber = registrationNumber
			pl.Slots[i].Color = color
			pl.CarSlotMap[registrationNumber] = i // Update map
			fmt.Printf("✅ Allocated slot number: %d\n", pl.Slots[i].SlotNumber)
			return
		}
	}
	fmt.Println("❌ Sorry, parking lot is full")
}

// Optimized Leave function dengan map lookup
func (pl *ParkingLot) Leave(registrationNumber string, hours int) {
	slotIndex, exists := pl.CarSlotMap[registrationNumber]
	if !exists {
		fmt.Printf("❌ Registration number %s not found\n", registrationNumber)
		return
	}

	slotNumber := pl.Slots[slotIndex].SlotNumber
	charge := calculateCharge(hours)
	
	// Free the slot
	pl.Slots[slotIndex].IsOccupied = false
	pl.Slots[slotIndex].RegistrationNumber = ""
	pl.Slots[slotIndex].Color = ""
	delete(pl.CarSlotMap, registrationNumber) // Remove from map
	
	fmt.Printf("🚗 Registration number %s with Slot Number %d is free with Charge $%d\n", 
		registrationNumber, slotNumber, charge)
}

// Fixed output format
func (pl *ParkingLot) Status() {
	fmt.Println("\n======= PARKING STATUS =======")
	fmt.Printf("%-8s %-15s\n", "Slot No.", "Registration No.")
	fmt.Println("------------------------------")
	hasOccupied := false
	for i := 0; i < pl.Capacity; i++ {
		if pl.Slots[i].IsOccupied {
			fmt.Printf("%-8d %-15s\n", pl.Slots[i].SlotNumber, pl.Slots[i].RegistrationNumber)
			hasOccupied = true
		}
	}
	if !hasOccupied {
		fmt.Println("No cars parked")
	}
	fmt.Println("==============================\n")
}

func calculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + (hours-2)*10
}

// Tambahan: Function untuk testing
func RunCommands(commands []string) {
	parkingLot := &ParkingLot{}
	for _, command := range commands {
		processCommand(parkingLot, command)
	}
}

func processCommand(pl *ParkingLot, command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create_parking_lot":
		if len(parts) >= 2 {
			capacity, err := strconv.Atoi(parts[1])
			if err == nil {
				pl.CreateParkingLot(capacity)
			}
		}
	case "park":
		// Updated untuk handle warna mobil
		if len(parts) >= 3 {
			pl.Park(parts[1], parts[2])
		} else if len(parts) >= 2 {
			// Fallback jika tidak ada warna (untuk backward compatibility)
			pl.Park(parts[1], "Unknown")
		}
	case "leave":
		if len(parts) >= 3 {
			hours, err := strconv.Atoi(parts[2])
			if err == nil {
				pl.Leave(parts[1], hours)
			}
		}
	case "status":
		pl.Status()
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("🚗 PARKING LOT MANAGEMENT SYSTEM 🚗")
	fmt.Println("===================================\n")

	parkingLot := &ParkingLot{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		if command != "" {
			processCommand(parkingLot, command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	fmt.Println("\n=== PROGRAM COMPLETED ===")
}