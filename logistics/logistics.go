package logistics
<<<<<<< HEAD
import "fmt"
=======

import (
	"fmt"
	"os"
)

>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
type Parcel struct {
	TrackingNumber    string
	WeightKG          float64
	DestinationPostal string
	Status            string
}
var parcels = make(map[string]Parcel)
<<<<<<< HEAD
=======

// Register Parcel
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
func RegisterParcel(p Parcel) error {

	if _, exists := parcels[p.TrackingNumber]; exists {
		return fmt.Errorf("parcel already exists")
	}
	p.Status = "Received"
	parcels[p.TrackingNumber] = p

	// Save parcel to parcel.txt
	file, err := os.OpenFile("parcel.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintf(file,
		"Tracking Number : %s\nWeight : %.2f KG\nDestination : %s\nStatus : %s\n--------------------------\n",
		p.TrackingNumber,
		p.WeightKG,
		p.DestinationPostal,
		p.Status,
	)

	if err != nil {
		return err
	}

	return nil
}

// Track Parcel
func TrackParcel(trackingNum string) (*Parcel, error) {

	p, exists := parcels[trackingNum]

	if !exists {
		return nil, fmt.Errorf("parcel not found")
	}
	return &p, nil
}
<<<<<<< HEAD
=======

// Update Status
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
func UpdateStatus(trackingNum string, newStatus string) error {

	p, exists := parcels[trackingNum]

	if !exists {
		return fmt.Errorf("parcel not found")
	}
	p.Status = newStatus
	parcels[trackingNum] = p

	return nil
}
<<<<<<< HEAD
=======

// Delete Delivered Parcel
>>>>>>> bdd023b (Added parcel.txt storage using command line arguments)
func PurgeParcelRecord(trackingNum string) error {

	p, exists := parcels[trackingNum]

	if !exists {
		return fmt.Errorf("parcel not found")
	}

	if p.Status != "Delivered" {
		return fmt.Errorf("parcel is not delivered")
	}

	delete(parcels, trackingNum)

	return nil

}
