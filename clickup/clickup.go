package clickup

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BASE_URL = "https://app.clickup/api/v2"
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

	baseURL, err := url.Parse(BASE_URL)

	if err != nil {
		return &Client{}, err
	}

	client := &Client{httpClient: httpClient, baseURL: baseURL}
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

// includeHeaders set HTTP headers from client.headers to *http.Request
func (c *Client) includeHeaders(req *http.Request) {
	for key, value := range c.headers {
		req.Header.Set(key, value)
	}
}

func (c *Client) prepareRequest(ctx context.Context, req *http.Request) *http.Request {
	out := req.WithContext(ctx)
	c.includeHeaders(out)
	if c.credential != "" {
		// throw an error
	}

	return out
}

// get get JSON data from API and returns its body as []bytes
func (c *Client) get(ctx context.Context, path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL.String()+path, nil)
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}
	return body, nil
}

func (c *Client) post(ctx context.Context, path string, data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}

	req = c.prepareRequest(ctx, req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// SetEndpointURL replace full URL of endpoint without subdomain validation.
// This is mainly used for testing to point to mock API server.
func (c *Client) SetEndpointURL(newURL string) error {
	baseURL, err := url.Parse(newURL)
	if err != nil {
		return err
	}

	c.baseURL = baseURL
	return nil
}
