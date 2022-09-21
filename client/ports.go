package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetRules(interf PortsReq) ([]NetRule, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/rules/%v", interf.Instance, interf.Vlan), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := []NetRule{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (c *Client) CreatePort(interf PortsReq, port OpenPort) (*OpenPort, error) {
	data, err := json.Marshal(port)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/port/%v/", interf.Instance, interf.Vlan), "POST", *bytes.NewBuffer(data), 201)
	if err != nil {
		return nil, err
	}
	rules := []NetRule{}
	err = json.NewDecoder(body).Decode(&rules)
	if err != nil {
		return nil, err
	}
	port.Forwarding = true // TODO: implement for WAR networks
	for _, r := range rules {
		if r.Dport == port.DestinationPort {
			port.SourcePort = r.NatExternalPort
		}
	}
	return &port, nil
}

func (c *Client) DeletePort(interf PortsReq, port OpenPort) error {
	data, err := json.Marshal(port)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/port/%v/", interf.Instance, interf.Vlan), "DELETE", *bytes.NewBuffer(data), 204)
	if err != nil {
		return err
	}
	return nil
}
