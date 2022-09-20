package provider

import (
	"context"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceGroupRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"users": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	if _, ok := d.GetOk("name"); ok {
		tflog.Info(ctx, "Get group by name")
		group, err := c.GetGroupByName(d.Get("name").(string))
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(strconv.Itoa(group.ID))
		d.Set("users", group.UserSet)
	} else if _, ok = d.GetOk("id"); ok {
		tflog.Info(ctx, "get group by id")
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		group, err := c.GetGroupByID(id)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.Itoa(group.ID))
		d.Set("name", group.Name)
		d.Set("users", group.UserSet)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One required field: name or id",
			Detail:   "",
		})
	}

	return diags
}
