package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	hostname   string
	port       int
	authToken  string
	httpClient *http.Client
}

func NewClient(hostname string, port int, token string) *Client {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	return &Client{
		hostname:   hostname,
		port:       port,
		authToken:  token,
		httpClient: &http.Client{Timeout: 10 * time.Second, Transport: transCfg},
	}
}

func (c *Client) GetAllLeases() ([]Lease, error) {
	body, err := c.httpRequest("dashboard/acpi/lease/", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	items := []Lease{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("token %s", c.authToken))
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
