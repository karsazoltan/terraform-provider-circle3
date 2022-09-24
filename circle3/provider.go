package provider

import (
	"terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
			},
			"port": {
				Type:        schema.TypeInt,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CIRCLE3_TOKEN", ""),
			},
			"datacenters": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("SERVICE_ADDRESS", ""),
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("SERVICE_PORT", ""),
						},
						"token": {
							Type:        schema.TypeString,
							Required:    true,
							DefaultFunc: schema.EnvDefaultFunc("CIRCLE3_TOKEN", ""),
						},
					},
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"circle3_vm":       resourceVM(),
			"circle3_disk":     resourceDisk(),
			"circle3_template": resourceTemplate(),
			"circle3_vmpool":   resourceVMPool(),
			"circle3_port":     resourcePort(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"circle3_leases":   dataSourceLeases(),
			"circle3_lease":    dataSourceLeasesByName(),
			"circle3_vlans":    dataSourceVlans(),
			"circle3_vlan":     dataSourceVlan(),
			"circle3_user":     dataSourceUser(),
			"circle3_group":    dataSourceGroup(),
			"circle3_template": dataSourceTemplate(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	if _, ok := d.GetOk("datacenters"); ok {
		datacenters_int := d.Get("datacenters").([]interface{})
		clients := make([]client.Client, len(datacenters_int))
		for i, e := range datacenters_int {
			item := e.(schema.ResourceData)
			address := item.Get("address").(string)
			port := item.Get("port").(int)
			token := item.Get("token").(string)
			clients[i] = *client.NewClient(address, port, token)
		}
	} else {
		address := d.Get("address").(string)
		port := d.Get("port").(int)
		token := d.Get("token").(string)
		return client.NewClient(address, port, token), nil
	}
	return nil, nil
}
