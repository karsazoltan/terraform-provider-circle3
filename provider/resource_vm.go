package provider

import (
	"context"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMUpdate,
		DeleteContext: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"vms": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"access_method": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"boot_menu": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"lease": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"cloud_init": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"ci_meta_data": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ci_user_data": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"system": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"has_agent": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"num_cores": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ram_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"max_ram_size": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"arch": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vms := d.Get("vms").([]interface{})
	vmis := []circleclient.Lease{}

	for _, item := range vms {
		i := item.(map[string]interface{})

		co := i["vm"].([]interface{})[0]
		coffee := co.(map[string]interface{})

		vm := hc.OrderItem{
			Coffee: hc.Coffee{
				ID: coffee["id"].(int),
			},
		}

		vms = append(vms, vm)
	}

	return diags
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceVMRead(ctx, d, m)
}

func resourceVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}
