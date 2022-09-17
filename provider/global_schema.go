package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func vmSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"pw": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"node": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"ipv4": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ipv6": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"owner": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"access_method": {
			Type:     schema.TypeString,
			Required: true,
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"boot_menu": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"lease": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"cloud_init": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"ci_meta_data": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ci_user_data": {
			Type:     schema.TypeString,
			Required: true,
		},
		"system": {
			Type:     schema.TypeString,
			Required: true,
		},
		"has_agent": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"num_cores": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"ram_size": {
			Type:             schema.TypeInt,
			Required:         true,
			ValidateDiagFunc: ValidateRamNumber,
		},
		"max_ram_size": {
			Type:             schema.TypeInt,
			Required:         true,
			ValidateDiagFunc: ValidateRamNumber,
		},
		"arch": {
			Type:     schema.TypeString,
			Required: true,
		},
		"priority": {
			Type:             schema.TypeInt,
			Required:         true,
			ValidateDiagFunc: ValidatePositiveNumber,
		},
		"vlans": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Optional: true,
		},
		"disks": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Optional: true,
		},
	}
}

func templateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"fromvm": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"owner": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"access_method": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"boot_menu": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"lease": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"cloud_init": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"ci_meta_data": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ci_user_data": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"system": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"has_agent": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"num_cores": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"ram_size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"max_ram_size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arch": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"priority": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"disks": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Computed: true,
		},
	}
}
