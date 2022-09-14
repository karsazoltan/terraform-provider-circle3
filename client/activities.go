package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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
