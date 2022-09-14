package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetInstanceActivities(id int) (*InstanceActivities, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vmact/%v/", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := InstanceActivities{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) GetStorageActivity(id int) (*StorageActivity, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/stact/%v/", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := StorageActivity{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
