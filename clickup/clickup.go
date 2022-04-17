package clickup

import (
	"net/http"
	"net/url"
)

const (
	BASE_URL = "https://app.clickup/api/v2/"
)

var defaultHeaders = map[string]string{
	"User-Agent":   "catdevman/go-clickup",
	"Content-Type": "application/json",
}

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	credential string
	headers    map[string]string
}

func NewClient(httpClient *http.Client) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{httpClient: httpClient}
	client.headers = defaultHeaders
	return client, nil
}

// SetHeader saves HTTP header in client. It will be included all API request
func (c *Client) SetHeader(key string, value string) {
	c.headers[key] = value
}

func (c *Client) SetCredential(pk string) {
	c.credential = pk
}
