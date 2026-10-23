package main

import (
	"fmt"
	"os"
	"strconv"

	"go_program/logistics"
)

func main() {
	// Check command-line arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <TrackingNumber> <WeightKG> <DestinationPostal>")
		return
	}

	tracking := os.Args[1]

	weight, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil || weight <= 0 {
		fmt.Println("Invalid Weight")
		return
	}

	destination := os.Args[3]

	// Create Parcel
	p := logistics.Parcel{
		TrackingNumber:    tracking,
		WeightKG:          weight,
		DestinationPostal: destination,
	}

	// Register Parcel
	err = logistics.RegisterParcel(p)
	if err != nil {
		fmt.Println("Registration Error:", err)
		return
	}

	fmt.Println("Parcel Registered Successfully")

	// Track Parcel
	parcel, err := logistics.TrackParcel(tracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nParcel Details")
	fmt.Println("Tracking Number :", parcel.TrackingNumber)
	fmt.Printf("Weight (KG)     : %.2f\n", parcel.WeightKG)
	fmt.Println("Destination     :", parcel.DestinationPostal)
	fmt.Println("Status          :", parcel.Status)

	// Update Status to In-Transit
	err = logistics.UpdateStatus(tracking, "In-Transit")
	if err != nil {
		fmt.Println(err)
		return
	}

	parcel, err = logistics.TrackParcel(tracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nAfter Update")
	fmt.Println("Status :", parcel.Status)

	// Update Status to Delivered
	err = logistics.UpdateStatus(tracking, "Delivered")
	if err != nil {
		fmt.Println(err)
		return
	}

	parcel, err = logistics.TrackParcel(tracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nAfter Delivery")
	fmt.Println("Status :", parcel.Status)

	// Delete Delivered Parcel
	err = logistics.PurgeParcelRecord(tracking)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nParcel record deleted successfully.")
}
