package main

import (
	"fmt"

	"go_program/logistics"
)

func main() {

	// Register a parcel
	p := logistics.Parcel{
		TrackingNumber:    "TRK1001",
		WeightKG:          4.5,
		DestinationPostal: "641602",
	}

	err := logistics.RegisterParcel(p)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Track parcel
	parcel, err := logistics.TrackParcel("TRK1001")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Tracking Number :", parcel.TrackingNumber)
	fmt.Println("Weight (KG)     :", parcel.WeightKG)
	fmt.Println("Destination     :", parcel.DestinationPostal)
	fmt.Println("Status          :", parcel.Status)

	// Update status
	err = logistics.UpdateStatus("TRK1001", "In-Transit")
	if err != nil {
		fmt.Println(err)
		return
	}

	parcel, _ = logistics.TrackParcel("TRK1001")
	fmt.Println("\nAfter Update")
	fmt.Println("Status :", parcel.Status)

	// Deliver parcel
	err = logistics.UpdateStatus("TRK1001", "Delivered")
	if err != nil {
		fmt.Println(err)
		return
	}

	parcel, _ = logistics.TrackParcel("TRK1001")
	fmt.Println("\nAfter Delivery")
	fmt.Println("Status :", parcel.Status)

	// Purge parcel record
	err = logistics.PurgeParcelRecord("TRK1001")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nParcel record deleted successfully.")
}
