package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetDisk(id int) (*Disk, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/disk/%v", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := Disk{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) CreateCDisk(cdisk CDisk) (*Disk, error) {
	reqvm, err := json.Marshal(cdisk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/createdisk/", cdisk.Instance), "POST", *bytes.NewBuffer(reqvm), 201)
	if err != nil {
		return nil, err
	}
	ret := Disk{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) CreateDDisk(ddiskk DDisk) (*Activities, error) {
	reqvm, err := json.Marshal(ddiskk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/downloaddisk/", ddiskk.Instance), "POST", *bytes.NewBuffer(reqvm), 201)
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

func (c *Client) DeleteDisk(instance_id int, disk_id int) error {
	reqstruct := struct {
		Disk int `json:"disk"`
	}{Disk: disk_id}
	json, err := json.Marshal(reqstruct)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/destroydisk/", instance_id), "DELETE", *bytes.NewBuffer(json), 204)
	return err
}