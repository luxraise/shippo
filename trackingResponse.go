package shippo

type TrackingResponse struct {
	Carrier        string `json:"carrier"`
	TrackingNumber string `json:"tracking_number"`

	AddressFrom Address `json:"addressFrom"`
	AddressTo   Address `json:"addressTo"`

	Transaction string `json:"transaction"`
	OriginalETA string `json:"original_eta"`
	ETA         string `json:"eta"`

	ServiceLevel NameToken `json:"servicelevel"`

	TrackingStatus  TrackingStatus   `json:"tracking_status"`
	TrackingHistory []TrackingStatus `json:"tracking_history"`
}
