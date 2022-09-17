package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetTemplate(id int) (*Template, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/template/%v", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := Template{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) GetTemplateByName(name string) (*Template, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/template?name=%s", name), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := Template{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) CreateTemplateFromVM(vmid int, name string) (*InstanceActivities, error) {
	reqstruct := struct {
		Name string `json:"string"`
	}{Name: name}
	reqjson, err := json.Marshal(reqstruct)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/saveastemplate/", vmid), "POST", *bytes.NewBuffer(reqjson), 201)
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

func (c *Client) DeleteTemplate(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/template/%v/", id), "DELETE", bytes.Buffer{}, 204)
	return err
}
