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
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: ValidateStatus,
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
			Optional: true,
		},
		"access_method": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"boot_menu": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"lease": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"cloud_init": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ci_meta_data": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"ci_user_data": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"ci_network_config": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"hookurl": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
		},
		"system": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"has_agent": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"num_cores": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"num_cores_max": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"ram_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			ValidateDiagFunc: ValidateRamNumber,
		},
		"max_ram_size": {
			Type:             schema.TypeInt,
			Optional:         true,
			ValidateDiagFunc: ValidateRamNumber,
		},
		"arch": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"priority": {
			Type:             schema.TypeInt,
			Optional:         true,
			ValidateDiagFunc: ValidatePriority,
		},
		"vlans": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Optional: true,
			Computed: true,
		},
		"disks": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Optional: true,
		},
		"from_template": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"sshportipv4": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"hostipv4": {
			Type:     schema.TypeString,
			Computed: true,
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
		"ci_network_config": {
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

func diskSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"size": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"filename": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"datastore": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"bus": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"base": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"dev_num": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"destroyed": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ci_disk": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"is_ready": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"size_format": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"url": {
			Type:     schema.TypeString,
			Optional: true,
			ForceNew: true,
		},
		"resize": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "",
			ForceNew: true,
		},
	}
}

func vmpoolSchema() map[string]*schema.Schema {
	vms := vmSchema()
	vms["name"].Required = false
	vms["name"].Computed = true
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"from_template": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"status": {
			Type:             schema.TypeString,
			Optional:         true,
			ValidateDiagFunc: ValidateStatus,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"users": {
			Type: schema.TypeList,
			Elem: &schema.Schema{
				Type: schema.TypeInt,
			},
			Required: true,
		},
		"vms": {
			Type:     schema.TypeList,
			Elem:     &schema.Resource{Schema: vms},
			Computed: true,
		},
	}
}

func portSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"forwarding": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"source_port": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vm": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"vlan": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Required: true,
		},
	}
}

func variableSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"value": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"url": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}
