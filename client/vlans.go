package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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