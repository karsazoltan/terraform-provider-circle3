package provider

import (
	"context"
	"strconv"
	"time"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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

	if _, ok := d.GetOk("name"); ok {
		tflog.Info(ctx, "Get vlan by name")
		reqvlan, err := c.GetVlanByName(d.Get("name").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.Itoa(reqvlan.ID))
		d.Set("comment", reqvlan.Comment)
		d.Set("description", reqvlan.Description)
		d.Set("domain", reqvlan.Domain)
		d.Set("vid", reqvlan.Vid)
	} else if _, ok := d.GetOk("id"); ok {
		tflog.Info(ctx, "Get vlan by id")
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		reqvlan, err := c.GetVlanByID(id)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.Itoa(reqvlan.ID))
		d.Set("comment", reqvlan.Comment)
		d.Set("description", reqvlan.Description)
		d.Set("domain", reqvlan.Domain)
		d.Set("vid", reqvlan.Vid)
		d.Set("name", reqvlan.Name)
	}

	return diags
}
