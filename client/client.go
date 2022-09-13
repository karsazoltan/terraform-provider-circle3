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
		httpClient: &http.Client{Timeout: 30 * time.Second, Transport: transCfg},
	}
}

func (c *Client) GetAllLeases() ([]Lease, error) {
	body, err := c.httpRequest("dashboard/acpi/lease/", "GET", bytes.Buffer{}, 200)
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

func (c *Client) GetAllVlans() ([]Vlan, error) {
	body, err := c.httpRequest("dashboard/acpi/vlan/", "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	items := []Vlan{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) GetLeasesByName(name string) (*Lease, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/lease?name=%s", name), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := Lease{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) GetVlanByName(name string) (*Vlan, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vlan?name=%s", name), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := Vlan{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) CreateVM(vm VM) (*VM, error) {
	reqvm, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/vm/", "POST", *bytes.NewBuffer(reqvm), 201)
	if err != nil {
		return nil, err
	}
	retvm := VM{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm, nil
}

func (c *Client) CreateDDisk(ddiskk DDisk) (*Activities, error) {
	reqvm, err := json.Marshal(ddiskk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/downloaddisk/", ddiskk.Instance), "POST", *bytes.NewBuffer(reqvm), 201)
	if err != nil {
		return nil, err
	}
	ret := Activities{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) GetActivity(id int) (*Activities, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vmact/%v/", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := Activities{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) DeleteVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/", id), "DELETE", bytes.Buffer{}, 401)
	return err
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer, allowedStaus int) (closer io.ReadCloser, err error) {
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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != allowedStaus {
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
