package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetOpenPorts(interf PortsReq) ([]OpenPort, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/vlan/%v", interf.Instance, interf.Vlan), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := []OpenPort{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) CreatePort(interf PortsReq, port OpenPort) (*OpenPort, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/vlan/%v", interf.Instance, interf.Vlan), "POST", bytes.Buffer{}, 201)
	if err != nil {
		return nil, err
	}
	item := OpenPort{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) DeletePort(interf PortsReq, port OpenPort) (*OpenPort, error) {
	data, err := json.Marshal(port)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/vlan/%v", interf.Instance, interf.Vlan), "POST", *bytes.NewBuffer(data), 201)
	if err != nil {
		return nil, err
	}
	item := OpenPort{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
