package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	hostname   string
	port       int
	authToken  string
	httpClient *http.Client
}

func NewClient(hostname string, port int, token string) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAllLeases() (*map[string]Lease, error) {
	body, err := c.httpRequest("dashboard/acpi/lease/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := map[string]Lease{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.authToken)
	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}
