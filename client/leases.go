package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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

func (c *Client) GetLeasesByID(id int) (*Lease, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/lease/%v", id), "GET", bytes.Buffer{}, 200)
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
