package shippo

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

var (
	// ErrEmptyAPIKey is returned when a Client is initialized with an empty API key
	ErrEmptyAPIKey = errors.New("invalid API key, cannot be empty")
)

/*
curl https://api.goshippo.com/shipments/  \
    -H "Authorization: ShippoToken <API_TOKEN>" \
    -H "Content-Type: application/json"  \
    -d '{...}'

*/

const (
	host = "https://api.goshippo.com"

	endpointTransactions    = "/transactions/"
	endpointCarrierAccounts = "/carrier_accounts/"
	endpointTracks          = "/tracks/%s/%s"
)

// New initializes and returns a new Stripe Client
func New(apiKey string) (client *Client, err error) {
	if len(apiKey) == 0 {
		err = ErrEmptyAPIKey
		return
	}

	var c Client
	if c.u, err = url.Parse(host); err != nil {
		return
	}

	c.apiKey = apiKey
	client = &c
	return
}

type Client struct {
	hc http.Client
	u  *url.URL

	apiKey string
}

func (c *Client) GetCarrierAccounts(carriers ...string) (accounts []CarrierAccount, err error) {
	q := url.Values{
		"carrier": carriers,
	}

	url := c.getURL(endpointCarrierAccounts, q)

	var resp getCarrierAccountsResponse
	if err = c.request("GET", url, nil, &resp); err != nil {
		return
	}

	accounts = resp.Results
	return
}

func (c *Client) CreateLabel(request CreateLabelRequest) (created CreateLabelResponse, err error) {
	url := c.getURL(endpointTransactions, nil)
	err = c.request("POST", url, &request, &created)
	return
}

func (c *Client) GetTracking(carrier, trackingNumber string) (resp TrackingResponse, err error) {
	endpoint := fmt.Sprintf(endpointTracks, carrier, trackingNumber)
	url := c.getURL(endpoint, nil)
	err = c.request("GET", url, nil, &resp)
	return
}

func (c *Client) request(method, url string, request, response interface{}) (err error) {
	var body io.Reader
	if body, err = getRequestBody(request); err != nil {
		err = fmt.Errorf("error getting request body: %v", err)
		return
	}

	var req *http.Request
	if req, err = http.NewRequest(method, url, body); err != nil {
		err = fmt.Errorf("error creating request: %v", err)
		return
	}

	req.Header.Set("Authorization", "ShippoToken "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	if resp, err = c.hc.Do(req); err != nil {
		err = fmt.Errorf("error performing request: %v", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200, 201:
		return handleResponse(resp.Body, response)
	case 400, 404:
		return handleError(resp.Body)
	default:
		return fmt.Errorf("Unexpected status code of: %d (url: <%s>, method: <%s>)", resp.StatusCode, url, method)
	}
}

func (c *Client) getURL(endpoint string, q url.Values) string {
	u := *c.u
	u.Path = path.Join(endpoint)

	if q != nil {
		u.RawQuery = q.Encode()
	}

	return u.String()
}
