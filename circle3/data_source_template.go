package provider

import (
	"context"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceTemplateRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO template schema merging
		},
	}
}

func dataSourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	if _, ok := d.GetOk("name"); ok {
		tflog.Info(ctx, "Get template by name")
		template, err := c.GetTemplateByName(d.Get("name").(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("owner", template.Owner)
		d.Set("description", template.Description)
		d.Set("parent", template.Parent)
		d.Set("access_method", template.AccessMethod)
		d.Set("boot_menu", template.BootMenu)
		d.Set("lease", template.Lease)
		d.Set("raw_data", template.RawData)
		d.Set("cloud_init", template.CloudInit)
		d.Set("ci_meta_data", template.CiMetaData)
		d.Set("ci_user_data", template.CiUserData)
		d.Set("system", template.System)
		d.Set("has_agent", template.HasAgent)
		d.Set("num_cores", template.NumCores)
		d.Set("ram_size", template.RamSize)
		d.Set("max_ram_size", template.MaxRamSize)
		d.Set("arch", template.Arch)
		d.Set("priority", template.Priority)
		d.Set("disks", template.Disks)
	} else if _, ok := d.GetOk("id"); ok {
		tflog.Info(ctx, "Get lease by id")
		TemplateID, err := strconv.Atoi(d.Id())
		template, err := c.GetTemplate(TemplateID)
		if err != nil {
			return diag.FromErr(err)
		}
		d.Set("owner", template.Owner)
		d.Set("description", template.Description)
		d.Set("parent", template.Parent)
		d.Set("access_method", template.AccessMethod)
		d.Set("boot_menu", template.BootMenu)
		d.Set("lease", template.Lease)
		d.Set("raw_data", template.RawData)
		d.Set("cloud_init", template.CloudInit)
		d.Set("ci_meta_data", template.CiMetaData)
		d.Set("ci_user_data", template.CiUserData)
		d.Set("system", template.System)
		d.Set("has_agent", template.HasAgent)
		d.Set("num_cores", template.NumCores)
		d.Set("ram_size", template.RamSize)
		d.Set("max_ram_size", template.MaxRamSize)
		d.Set("arch", template.Arch)
		d.Set("priority", template.Priority)
		d.Set("disks", template.Disks)
	}
	return diags
}
