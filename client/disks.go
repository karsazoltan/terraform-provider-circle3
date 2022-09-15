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

func (c *Client) CreateDDisk(ddiskk DDisk) (*InstanceActivities, error) {
	reqvm, err := json.Marshal(ddiskk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/downloaddisk/", ddiskk.Instance), "POST", *bytes.NewBuffer(reqvm), 201)
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

func (c *Client) CreatePersistentCDisk(cdisk CDisk) (*Disk, error) {
	reqdisk, err := json.Marshal(cdisk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/pcdisk/", "POST", *bytes.NewBuffer(reqdisk), 201)
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

func (c *Client) CreatePersistentDDisk(ddiskk DDisk) (*StorageActivity, error) {
	reqdisk, err := json.Marshal(ddiskk)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/pddisk/", "POST", *bytes.NewBuffer(reqdisk), 201)
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

func (c *Client) DeletePersistentDisk(disk_id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/disk/%v", disk_id), "DELETE", bytes.Buffer{}, 204)
	return err
}

func (c *Client) AddNewPersistentDiskToVM(vm_id int, disk_id int) (*Disk, error) {
	reqstruct := struct {
		Disk int `json:"disk"`
	}{Disk: disk_id}
	data, err := json.Marshal(reqstruct)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/%v/addpersistent/", vm_id), "POST", *bytes.NewBuffer(data), 200)
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
