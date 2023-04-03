package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) DeleteLBVM(id int, datacenter string) error {
	q := url.Values{}
	q.Add("datacenter", datacenter)
	_, err := c.httpRequest(fmt.Sprintf("lb/dashboard/acpi/vm/%v/?%s", id, q.Encode()), "DELETE", bytes.Buffer{}, 204)
	return err
}

func (c *Client) GetLBVM(id int, datacenter string) (*LBVM, error) {
	q := url.Values{}
	q.Add("datacenter", datacenter)
	body, err := c.httpRequest(fmt.Sprintf("lb/dashboard/acpi/vm/%v?%s", id, q.Encode()), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := LBVM{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) CreateLBVMfromTemplate(template_name string, name string, owner string, balancermethod string) (*LBVM, error) {
	template := struct {
		TemplateName string `json:"template_name"`
		Name         string `json:"name"`
		Username     string `json:"username"`
	}{TemplateName: template_name, Name: name, Username: owner}
	req, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("lb/%s/dashboard/acpi/bvm/", balancermethod), "POST", *bytes.NewBuffer(req), 201)
	if err != nil {
		return nil, err
	}
	retvm := []LBVM{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm[0], nil
}
