package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) GetAllVariables() ([]Variable, error) {
	body, err := c.httpRequest("dashboard/acpi/var/", "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	items := []Variable{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) GetVariableByName(name string) (*Variable, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/var?key=%s", url.QueryEscape(name)), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := Variable{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) GetVariablesByID(id int) (*Variable, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/var/%v", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := Variable{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) CreateVariable(variable Variable) (*Variable, error) {
	reqvar, err := json.Marshal(variable)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/var/", "POST", *bytes.NewBuffer(reqvar), 201)
	if err != nil {
		return nil, err
	}
	item := Variable{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (c *Client) DeleteVariable(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/var/%v/", id), "DELETE", bytes.Buffer{}, 204)
	return err
}
