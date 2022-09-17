package provider

import (
	"context"
	"strconv"
	"time"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVlans() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVlansRead,
		Schema: map[string]*schema.Schema{
			"vlans": {
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
						"vid": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"domain": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"comment": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVlansRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	reqlease, err := c.GetAllVlans()
	if err != nil {
		return diag.FromErr(err)
	}

	leases := make([]map[string]interface{}, 0)

	for _, v := range reqlease {
		ingredient := make(map[string]interface{})

		ingredient["id"] = v.ID
		ingredient["vid"] = v.Vid
		ingredient["name"] = v.Name
		ingredient["comment"] = v.Comment
		ingredient["domain"] = v.Domain
		ingredient["description"] = v.Description

		leases = append(leases, ingredient)
	}

	if err := d.Set("vlans", leases); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func dataSourceVlanByName() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVlanByNameRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vid": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"domain": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVlanByNameRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	reqlease, err := c.GetVlanByName(d.Get("name").(string))
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(reqlease.ID))
	d.Set("comment", reqlease.Comment)
	d.Set("description", reqlease.Description)
	d.Set("domain", reqlease.Domain)
	d.Set("vid", reqlease.Vid)

	return diags
}
