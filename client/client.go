package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.medzakupivli.com",
	Path:   "/covid19/",
}

// Client is a client for sending requests to the MOZ API.
type Client struct {
	client *http.Client
}

// New creates a new Go client for the MOZ API.
func New() *Client {
	c := &http.Client{Timeout: 30 * time.Second}

	return &Client{
		client: c,
	}
}

// GetHostitals gets all list covid hospital.
func (c *Client) GetHostitals() (*Hostitals, error) {
	// set up a request to the hourly forecast endpoint
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "hospital_list"})

	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)

	if err != nil {
		return nil, err
	}

	// add URL headers, query params, then send the request
	req.Header.Add("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	// deserialize the response and return our hospital data
	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var hospitals Hostitals
		if err := json.NewDecoder(res.Body).Decode(&hospitals); err != nil {
			return nil, err
		}
		return &hospitals, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetProtective /covid19/protective
func (c *Client) GetProtective(args CovidArgs) (*Protectives, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "protective"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var protectives Protectives
		if err := json.NewDecoder(res.Body).Decode(&protectives); err != nil {
			return nil, err
		}
		return &protectives, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetEquipment /covid19/equipment
func (c *Client) GetEquipment(args CovidArgs) (*Equipment, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "equipment"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var equipment Equipment
		if err := json.NewDecoder(res.Body).Decode(&equipment); err != nil {
			return nil, err
		}
		return &equipment, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetMedicines /covid19/medicines
func (c *Client) GetMedicines(args CovidArgs) (*Medicines, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "medicines"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var medicines Medicines
		if err := json.NewDecoder(res.Body).Decode(&medicines); err != nil {
			return nil, err
		}
		return &medicines, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetBeds /covid19/beds
func (c *Client) GetBeds(args CovidArgs) (*Beds, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "beds"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var beds Beds
		if err := json.NewDecoder(res.Body).Decode(&beds); err != nil {
			return nil, err
		}
		return &beds, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		// handle unexpected status codes
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetPatients /covid19/patients
func (c *Client) GetPatients(args CovidArgs) (*Patients, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "patients"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var patients Patients
		if err := json.NewDecoder(res.Body).Decode(&patients); err != nil {
			return nil, err
		}
		return &patients, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}

// GetStaff /covid19/staff
func (c *Client) GetStaff(args CovidArgs) (*Staff, error) {
	endpt := baseURL.ResolveReference(
		&url.URL{Path: "staff"})
	req, err := http.NewRequest(http.MethodGet, endpt.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = args.QueryParams().Encode()

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case 200:
		var staff Staff
		if err := json.NewDecoder(res.Body).Decode(&staff); err != nil {
			return nil, err
		}
		return &staff, nil
	case 400, 401, 403, 500:
		//TODO Dry
		var errRes ErrorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return nil, err
		}

		if errRes.StatusCode == 0 {
			errRes.StatusCode = res.StatusCode
		}
		return nil, &errRes
	default:
		return nil, fmt.Errorf("unexpected status code %d", res.StatusCode)
	}
}
