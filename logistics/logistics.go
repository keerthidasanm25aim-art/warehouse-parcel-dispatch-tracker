package logistics

import (
	"fmt"
	"os"
)

type Parcel struct {
	TrackingNumber    string
	WeightKG          float64
	DestinationPostal string
	Status            string
}

var parcels = make(map[string]Parcel)

// Register Parcel
func RegisterParcel(p Parcel) error {
	if p.TrackingNumber == "" {
		return fmt.Errorf("tracking number cannot be empty")
	}

	if _, exists := parcels[p.TrackingNumber]; exists {
		return fmt.Errorf("parcel already exists")
	}

	p.Status = "Received"

	file, err := os.OpenFile(
		"parcel.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(
		file,
		"Tracking Number : %s\nWeight : %.2f KG\nDestination : %s\nStatus : %s\n--------------------------\n",
		p.TrackingNumber,
		p.WeightKG,
		p.DestinationPostal,
		p.Status,
	)

	closeErr := file.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}

	parcels[p.TrackingNumber] = p
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

// Update Status
func UpdateStatus(trackingNum string, newStatus string) error {
	p, exists := parcels[trackingNum]
	if !exists {
		return fmt.Errorf("parcel not found")
	}

	p.Status = newStatus
	parcels[trackingNum] = p

	return nil
}

// Delete Delivered Parcel
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
