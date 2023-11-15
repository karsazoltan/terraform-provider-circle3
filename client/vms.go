package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) DeleteVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/", id), "DELETE", bytes.Buffer{}, 204)
	return err
}

func (c *Client) DeleteVMs(ids []int) error {
	for _, id := range ids {
		_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/", id), "DELETE", bytes.Buffer{}, 204)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) GetVM(id int) (*VM, error) {
	body, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v", id), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := VM{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) UpdateVMResource(id int, resource VMResource) error {
	reqres, err := json.Marshal(resource)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%d/updateresource", id), "PUT", *bytes.NewBuffer(reqres), 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateMemHotplug(id int, mem_size int) error {
	mem := struct {
		RamSize int `json:"ram_size"`
	}{
		RamSize: mem_size,
	}
	reqres, err := json.Marshal(mem)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%d/hotplugmem/", id), "PUT", *bytes.NewBuffer(reqres), 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateVCPUHotplug(id int, num_vcpu int) error {
	vcpu := struct {
		NumCPU int `json:"num_cores"`
	}{
		NumCPU: num_vcpu,
	}
	reqres, err := json.Marshal(vcpu)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%d/hotplugvcpu/", id), "PUT", *bytes.NewBuffer(reqres), 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateVMLease(id int, lease_new int) error {
	lease := struct {
		Lease int `json:"lease"`
	}{
		Lease: lease_new,
	}
	reqres, err := json.Marshal(lease)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%d/updatelease", id), "PUT", *bytes.NewBuffer(reqres), 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeployVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/deploy/", id), "POST", bytes.Buffer{}, 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ShutdownVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/shutdown/", id), "POST", bytes.Buffer{}, 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SleepVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/sleep/", id), "POST", bytes.Buffer{}, 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) WakeUpVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/wakeup/", id), "POST", bytes.Buffer{}, 201)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateVMState(id int, old_state string, new_state string) error {
	states := map[string]map[string]interface{}{
		"STOPPED": {
			"RUNNING": c.DeployVM,
		},
		"PENDING": {
			"RUNNING": c.DeployVM,
		},
		"RUNNING": {
			"STOPPED":   c.ShutdownVM,
			"SUSPENDED": c.SleepVM,
		},
		"SUSPENDED": {
			"RUNNING": c.WakeUpVM,
		},
	}
	return states[old_state][new_state].(func(int) error)(id)
}

func (c *Client) CreateVM(vm VM) (*VM, error) {
	reqvm, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/vm/", "POST", *bytes.NewBuffer(reqvm), 201)
	if err != nil {
		return nil, err
	}
	retvm := VM{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm, nil
}

func (c *Client) CreateVMfromTemplate(template_id int, name string) (*VM, error) {
	template := struct {
		TemplateID int    `json:"template"`
		Name       string `json:"name"`
	}{TemplateID: template_id, Name: name}
	req, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/ft/", "POST", *bytes.NewBuffer(req), 201)
	if err != nil {
		return nil, err
	}
	retvm := []VM{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm[0], nil
}

func (c *Client) CreateVMfromRP(rpname string, key string) (*VM, error) {
	rp := struct {
		RPName string `json:"rpname"`
		Key    string `json:"key"`
	}{RPName: rpname, Key: key}
	req, err := json.Marshal(rp)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("rp/", "POST", *bytes.NewBuffer(req), 201)
	if err != nil {
		return nil, err
	}
	retvm := []VM{}
	err = json.NewDecoder(body).Decode(&retvm)
	if err != nil {
		return nil, err
	}
	return &retvm[0], nil
}

func (c *Client) CreateVMfromTemplateforUsers(template_id int, name string, users []int) ([]VM, error) {
	template := struct {
		TemplateID int    `json:"template"`
		Users      []int  `json:"users"`
		Name       string `json:"name"`
	}{TemplateID: template_id, Users: users, Name: name}
	req, err := json.Marshal(template)
	if err != nil {
		return nil, err
	}
	body, err := c.httpRequest("dashboard/acpi/ftusersid/", "POST", *bytes.NewBuffer(req), 201)
	if err != nil {
		return nil, err
	}
	retvms := []VM{}
	err = json.NewDecoder(body).Decode(&retvms)
	if err != nil {
		return nil, err
	}
	return retvms, nil
}

func (c *Client) GetVMPool(ids []int) ([]VM, error) {
	vms := make([]VM, 0)
	for _, e := range ids {
		vm, err := c.GetVM(e)
		if err != nil {
			return nil, err
		}
		vms = append(vms, *vm)
	}
	return vms, nil
}
