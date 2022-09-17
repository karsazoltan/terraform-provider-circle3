package provider

import (
	"context"
	"strconv"
	"time"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLeases() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLeasesRead,
		Schema: map[string]*schema.Schema{
			"leases": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"suspend_interval_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"delete_interval_seconds": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLeasesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)

	var diags diag.Diagnostics

	reqlease, err := c.GetAllLeases()
	if err != nil {
		return diag.FromErr(err)
	}

	leases := make([]map[string]interface{}, 0)

	for _, v := range reqlease {
		ingredient := make(map[string]interface{})

		ingredient["id"] = v.ID
		ingredient["name"] = v.Name
		ingredient["delete_interval_seconds"] = v.Delete_interval_seconds
		ingredient["suspend_interval_seconds"] = v.Suspend_interval_seconds

		leases = append(leases, ingredient)
	}

	if err := d.Set("leases", leases); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceLeasesByName() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLeasesByNameRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"suspend_interval_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"delete_interval_seconds": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceLeasesByNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	reqlease, err := c.GetLeasesByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(reqlease.ID))
	d.Set("delete_interval_seconds", reqlease.Delete_interval_seconds)
	d.Set("suspend_interval_seconds", reqlease.Suspend_interval_seconds)

	return diags
}
