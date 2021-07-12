package shippo

type Rate struct {
	ObjectID          string `json:"object_id"`
	Amount            string `json:"amount"`
	Currency          string `json:"currency"`
	AmountLocal       string `json:"amount_local"`
	CurrencyLocal     string `json:"currency_local"`
	Provider          string `json:"provider"`
	ServicelevelName  string `json:"servicelevel_name"`
	ServicelevelToken string `json:"servicelevel_token"`
	CarrierAccount    string `json:"carrier_account"`
}
