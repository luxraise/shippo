package shippo

type getCarrierAccountsResponse struct {
	Results []CarrierAccount `json:"results"`
}

type CarrierAccount struct {
	ObjectMeta

	Carrier   string `json:"carrier"`
	AccountID string `json:"account_id"`

	Active bool `json:"active"`
	Test   bool `json:"test"`

	IsShippoAccount bool   `json:"is_shippo_account"`
	Metadata        string `json:"metadata"`

	Parameters map[string]interface{} `json:"parameters"`
}
