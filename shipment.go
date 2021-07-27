package shippo

type Shipment struct {
	AddressFrom Address  `json:"address_from"`
	AddressTo   Address  `json:"address_to" form:"parcels"`
	Parcels     []Parcel `json:"parcels"`
}
