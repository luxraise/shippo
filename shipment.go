package shippo

// MakeShipment is a quality of life helper function to make a new instance of Shipment
func MakeShipment(from, to Address, parcels ...Parcel) (s Shipment) {
	s.AddressFrom = from
	s.AddressTo = to
	s.Parcels = parcels
	return
}

type Shipment struct {
	AddressFrom Address  `json:"address_from" form:"address_form"`
	AddressTo   Address  `json:"address_to" form:"address_to"`
	Parcels     []Parcel `json:"parcels" form:"parcels"`
}
