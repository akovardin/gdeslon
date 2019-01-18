package gdeslon

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type ApiError struct {
	ErrorName        string `json:"error"`
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e ApiError) Error() string {
	return e.ErrorName + ": " + e.ErrorDescription
}

type Client struct {
	key string

	url    string
	client *http.Client
}

func NewClient(url, key string) *Client {
	return &Client{
		url:    url,
		key:    key,
		client: &http.Client{},
	}
}
func (c *Client) Call(url, method string, params url.Values, result interface{}) error {
	response, err := c.request(url, method, params)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		return err
	}

	return err
}

func (c *Client) request(u, m string, params url.Values) (*http.Response, error) {
	params.Add("_gs_at", c.key)
	u = c.url + "/" + u + "/?" + params.Encode()
	request, err := http.NewRequest(m, u, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		return response, nil
	}

	defer response.Body.Close()

	e := ApiError{}
	if err := json.NewDecoder(response.Body).Decode(&e); err != nil {
		return nil, err
	}
	return nil, e
}
