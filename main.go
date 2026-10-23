package main
import (
	"fmt"
<<<<<<< HEAD
=======
	"os"
	"strconv"

>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	"go_program/logistics"
)
func main() {
<<<<<<< HEAD
	p := logistics.Parcel{
		TrackingNumber:    "TRK1001",
		WeightKG:          4.5,
		DestinationPostal: "641602",
	}
	err := logistics.RegisterParcel(p)
=======

	// Check command-line arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . <TrackingNumber> <WeightKG> <DestinationPostal>")
		return
	}

	tracking := os.Args[1]

	weight, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
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
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
<<<<<<< HEAD
	parcel, err := logistics.TrackParcel("TRK1001")
=======

	// Track Parcel
	parcel, err := logistics.TrackParcel(tracking)
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Tracking Number :", parcel.TrackingNumber)
	fmt.Println("Weight (KG)     :", parcel.WeightKG)
	fmt.Println("Destination     :", parcel.DestinationPostal)
	fmt.Println("Status          :", parcel.Status)
<<<<<<< HEAD
	err = logistics.UpdateStatus("TRK1001", "In-Transit")
=======

	// Save Input to File (Append Mode)
	file, err := os.OpenFile("parcel.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Fprintln(file, "Parcel Details")
	fmt.Fprintln(file, "Tracking Number :", parcel.TrackingNumber)
	fmt.Fprintln(file, "Weight (KG)     :", parcel.WeightKG)
	fmt.Fprintln(file, "Destination     :", parcel.DestinationPostal)
	fmt.Fprintln(file, "Status          :", parcel.Status)
	fmt.Fprintln(file, "------------------------------------")

	fmt.Println("\nInput saved to parcel.txt")

	// Update Status
	err = logistics.UpdateStatus(tracking, "In-Transit")
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
<<<<<<< HEAD
	parcel, _ = logistics.TrackParcel("TRK1001")
	fmt.Println("\nAfter Update")
	fmt.Println("Status :", parcel.Status)
	err = logistics.UpdateStatus("TRK1001", "Delivered")
=======

	parcel, _ = logistics.TrackParcel(tracking)

	fmt.Println("\nAfter Update")
	fmt.Println("Status :", parcel.Status)

	// Deliver Parcel
	err = logistics.UpdateStatus(tracking, "Delivered")
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
<<<<<<< HEAD
	parcel, _ = logistics.TrackParcel("TRK1001")
	fmt.Println("\nAfter Delivery")
	fmt.Println("Status :", parcel.Status)
	err = logistics.PurgeParcelRecord("TRK1001")
=======

	parcel, _ = logistics.TrackParcel(tracking)

	fmt.Println("\nAfter Delivery")
	fmt.Println("Status :", parcel.Status)

	// Delete Record
	err = logistics.PurgeParcelRecord(tracking)
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nParcel record deleted successfully.")
}
