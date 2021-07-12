# Shippo
Shippo is a Client SDK for the Shippo API

## Usage
### New
```go
func ExampleNew() {
	var err error
	if testClient, err = New("[Shippo API Key]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shippo Client has been initialized! %v\n", testClient)
}

```

### Client.GetCarrierAccounts
```go
func ExampleClient_GetCarrierAccounts() {
	var (
		accounts []CarrierAccount
		err      error
	)

	if accounts, err = testClient.GetCarrierAccounts("usps"); err != nil {
		log.Fatal(err)
	}

	if len(accounts) == 0 {
		log.Fatal("invalid number of carrier accounts, expected at least one and received none")
	}

	// Get the first carrier account ID
	testCarrierAccountID = accounts[0].ObjectID
}
```

### Client.CreateLabel
```go
func ExampleClient_CreateLabel() {
	var (
		response CreateLabelResponse
		err      error
	)

	var request CreateLabelRequest
	request.CarrierAccount = testCarrierAccountID
	request.ServicelevelToken = "usps_priority"

	request.Shipment.AddressFrom.Street1 = "614 SW 11th Ave"
	request.Shipment.AddressFrom.City = "Portland"
	request.Shipment.AddressFrom.State = "OR"
	request.Shipment.AddressFrom.Zip = "97205"
	request.Shipment.AddressFrom.Country = "USA"
	request.Shipment.AddressFrom.Name = "Jackknife Bar"

	request.Shipment.AddressTo.Street1 = "7000 NE Airport Way"
	request.Shipment.AddressTo.City = "Portland"
	request.Shipment.AddressTo.State = "OR"
	request.Shipment.AddressTo.Zip = "97218"
	request.Shipment.AddressTo.Country = "USA"
	request.Shipment.AddressTo.Name = "Portland Airport"

	request.Shipment.Parcels = []Parcel{
		{
			Length:       "5",
			Width:        "5",
			Height:       "5",
			DistanceUnit: "in",
			Weight:       "2",
			MassUnit:     "lb",
		},
	}

	if response, err = testClient.CreateLabel(request); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shippo response: %+v", response)
}
```
