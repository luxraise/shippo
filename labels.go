package shippo

import "time"

type CreateLabelRequest struct {
	Shipment Shipment `json:"shipment"`
	// Account ID for requested carrier
	CarrierAccount string `json:"carrier_account"`
	// Service level for parcel to be shipped with (e.g. usps_priority)
	ServicelevelToken string `json:"servicelevel_token"`
}

type CreateLabelResponse struct {
	ObjectMeta

	Status string `json:"status"`

	TrackingNumber      string               `json:"tracking_number"`
	TrackingStatus      *MaybeTrackingStatus `json:"tracking_status"`
	TrackingURLProvider string               `json:"tracking_url_provider"`

	Rate Rate      `json:"rate"`
	Eta  time.Time `json:"eta"`

	Metadata string        `json:"metadata"`
	Messages []interface{} `json:"messages"`

	LabelURL             string `json:"label_url"`
	CommercialInvoiceURL string `json:"commercial_invoice_url"`

	WasTest bool `json:"was_test"`
}
