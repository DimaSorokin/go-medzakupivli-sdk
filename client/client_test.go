package client

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func TestClient_GetHostitals(t *testing.T) {
	//json := `{"name":"Test Name","full_name":"test full name","owner":{"login": "octocat"}}`
	//r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://api.medzakupivli.com/covid19/hospital_list")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
			Header:     make(http.Header),
		}
	})

	api := Client{client: client}
	body, err := api.GetHostitals()
	assert.Error(t, err)
	assert.NotEqual(t, []byte("OK"), body)
}
