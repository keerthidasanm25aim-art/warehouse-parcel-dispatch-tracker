package logistics
import "fmt"
type Parcel struct {
	TrackingNumber    string
	WeightKG          float64
	DestinationPostal string
	Status            string
}
var parcels = make(map[string]Parcel)
func RegisterParcel(p Parcel) error {
	if _, exists := parcels[p.TrackingNumber]; exists {
		return fmt.Errorf("parcel already exists")
	}
	p.Status = "Received"
	parcels[p.TrackingNumber] = p
	return nil
}
func TrackParcel(trackingNum string) (*Parcel, error) {
	p, exists := parcels[trackingNum]
	if !exists {
		return nil, fmt.Errorf("parcel not found")
	}
	return &p, nil
}
func UpdateStatus(trackingNum string, newStatus string) error {
	p, exists := parcels[trackingNum]
	if !exists {
		return fmt.Errorf("parcel not found")
	}
	p.Status = newStatus
	parcels[trackingNum] = p
	return nil
}
func PurgeParcelRecord(trackingNum string) error {
	p, exists := parcels[trackingNum]
	if !exists {
		return fmt.Errorf("parcel not found")
	}

	if p.Status != "Delivered" {
		return fmt.Errorf("parcel is not delivered yet")
	}

	delete(parcels, trackingNum)
	return nil
}
