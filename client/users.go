package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
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

func (c *Client) GetUserByName(username string) (*User, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/user?username=%s", url.QueryEscape(username)), "GET", bytes.Buffer{}, 200)
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

func (c *Client) GetUserByID(id int) (*User, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/user/%v", id), "GET", bytes.Buffer{}, 200)
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
