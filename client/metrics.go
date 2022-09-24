package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetMetric(target string, deltatime_h int) ([]Metric, error) {
	body, err := c.httpRequest(fmt.Sprintf(".usage2?target%s&time=%vh", target, deltatime_h), "GET", bytes.Buffer{}, 200)
	if err != nil {
		return nil, err
	}
	ret := []Metric{}
	err = json.NewDecoder(body).Decode(&ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (c *Client) GetAVGCPU(deltatime_h int) ([]Metric, error) {
	return c.GetMetric("avg(circle.bm*.cpu.percent)", deltatime_h)
}

func (c *Client) GetAVGMem(deltatime_h int) ([]Metric, error) {
	return c.GetMetric("avg(circle.bm*.memory.usage)", deltatime_h)
}

func (c *Client) GetSumVM(deltatime_h int) ([]Metric, error) {
	return c.GetMetric("sum(circle.bm*.vmcount)", deltatime_h)
}
