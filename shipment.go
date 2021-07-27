package shippo

type Shipment struct {
	AddressFrom Address  `json:"address_from" form:"address_form"`
	AddressTo   Address  `json:"address_to" form:"address_to"`
	Parcels     []Parcel `json:"parcels" form:"parcels"`
}
