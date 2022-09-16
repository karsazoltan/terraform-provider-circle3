package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllUsers() ([]User, error) {
	body, err := c.httpRequest("dashboard/acpi/user/", "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	items := []User{}
	err = json.NewDecoder(body).Decode(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *Client) GetUserByName(name string) (*User, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/user?name=%s", name), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	item := User{}
	err = json.NewDecoder(body).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
