package provider

import (
	"context"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resource_resourceGroups() *schema.Resource {
	return &schema.Resource{
		CreateContext: resource_resourceGroupsCreate,
		ReadContext:   resource_resourceGroupsRead,
		UpdateContext: resource_resourceGroupsUpdate,
		DeleteContext: resource_resourceGroupsDelete,
		Schema:        rpSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resource_resourceGroupsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*circleclient.Client)

	if _, ok := d.GetOk("num_vms"); ok {
		resp, err := c.CreateRPLoadBalancing(
			d.Get("rpname").(string), d.Get("from_template").(string),
			d.Get("num_vms").(int), d.Get("key").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(resp.ID)
	}

	return diags
}

func resource_resourceGroupsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resource_resourceGroupsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resource_resourceGroupsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
