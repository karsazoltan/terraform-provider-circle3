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

func (c *Client) DeployVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/deploy/", id), "POST", bytes.Buffer{}, 200)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ShutdownVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/shutdown/", id), "POST", bytes.Buffer{}, 200)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SleepVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/sleep/", id), "POST", bytes.Buffer{}, 200)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) WakeUpVM(id int) error {
	_, err := c.httpRequest(fmt.Sprintf("dashboard/acpi/vm/%v/wakeup/", id), "POST", bytes.Buffer{}, 200)
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