package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllGroups() ([]Group, error) {
	body, err := c.httpRequest("dashboard/acpi/group/", "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	items := []Group{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) GetGroupByName(name string) (*Group, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/group?name=%s", name), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := Group{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
