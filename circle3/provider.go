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
		},
		ResourcesMap: map[string]*schema.Resource{
			"circle3_vm":       resourceVM(),
			"circle3_disk":     resourceDisk(),
			"circle3_template": resourceTemplate(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"circle3_leases":          dataSourceLeases(),
			"circle3_lease_byname":    dataSourceLeasesByName(),
			"circle3_vlans":           dataSourceVlans(),
			"circle3_vlan_byname":     dataSourceVlanByName(),
			"circle3_user_byusername": dataSourceUserByUsername(),
			"circle3_group_byname":    dataSourceGroupByName(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	address := d.Get("address").(string)
	port := d.Get("port").(int)
	token := d.Get("token").(string)
	return client.NewClient(address, port, token), nil
}
