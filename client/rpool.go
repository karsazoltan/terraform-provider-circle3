package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) CreateRPLoadBalancing(rp_name string, template_name string, num_vms int, key string) (*RP, error) {
	rp := struct {
		RPName       string `json:"rpname"`
		FromTemplate string `json:"from_template"`
		NumVMs       int    `json:"num_vms"`
		Key          string `json:"key"`
	}{RPName: rp_name, FromTemplate: template_name, NumVMs: num_vms}
	req, err := json.Marshal(rp)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("createrp/", "POST", *bytes.NewBuffer(req), 201)
	if err != nil {
		return nil, err
	}
	retvm := []RP{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm[0], nil
}
