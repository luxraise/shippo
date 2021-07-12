package shippo

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	testAPIKey           = os.Getenv("SHIPPO_TEST_API_KEY")
	testClient           *Client
	testCarrierAccountID string
)

func TestClient_New(t *testing.T) {
	var err error
	type testcase struct {
		apiKey string
		want   error
	}

	tcs := []testcase{
		{
			apiKey: testAPIKey,
			want:   nil,
		},
		{
			apiKey: "",
			want:   ErrEmptyAPIKey,
		},
	}

	for _, tc := range tcs {
		if _, err = New(tc.apiKey); err != tc.want {
			t.Fatalf("invalid error, expected %v and recieved %v", tc.want, err)
		}
	}
}

func TestClient_GetCarrierAccounts(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	var accounts []CarrierAccount
	if accounts, err = c.GetCarrierAccounts("usps"); err != nil {
		t.Fatal(err)
	}

	if len(accounts) == 0 {
		t.Fatal("invalid number of carrier accounts, expected at least one and received none")
	}
}

func TestClient_CreateLabel(t *testing.T) {
	var (
		c   *Client
		err error
	)

	if c, err = New(testAPIKey); err != nil {
		t.Fatal(err)
	}

	var accounts []CarrierAccount
	if accounts, err = c.GetCarrierAccounts("usps"); err != nil {
		t.Fatal(err)
	}

	if len(accounts) == 0 {
		t.Fatal("invalid number of carrier accounts, expected at least one and received none")
	}

	var label CreateLabelRequest
	label.CarrierAccount = accounts[0].ObjectID
	label.ServicelevelToken = "usps_priority"

	label.Shipment.AddressFrom.Street1 = "614 SW 11th Ave"
	label.Shipment.AddressFrom.City = "Portland"
	label.Shipment.AddressFrom.State = "OR"
	label.Shipment.AddressFrom.Zip = "97205"
	label.Shipment.AddressFrom.Country = "USA"
	label.Shipment.AddressFrom.Name = "Jackknife Bar"

	label.Shipment.AddressTo.Street1 = "7000 NE Airport Way"
	label.Shipment.AddressTo.City = "Portland"
	label.Shipment.AddressTo.State = "OR"
	label.Shipment.AddressTo.Zip = "97218"
	label.Shipment.AddressTo.Country = "USA"
	label.Shipment.AddressTo.Name = "Portland Airport"

	label.Shipment.Parcels = []Parcel{
		{
			Length:       "5",
			Width:        "5",
			Height:       "5",
			DistanceUnit: "in",
			Weight:       "2",
			MassUnit:     "lb",
		},
	}

	if _, err = c.CreateLabel(label); err != nil {
		t.Fatal(err)
	}
}

func ExampleNew() {
	var err error
	if testClient, err = New("[Shippo API Key]"); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Shippo Client has been initialized! %v\n", testClient)
}

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
