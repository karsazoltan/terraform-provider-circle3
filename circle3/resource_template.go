package provider

import (
	"context"
	"fmt"
	"strconv"
	circleclient "terraform-provider-circle3/client"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceTemplateCreate,
		ReadContext:   resourceTemplateRead,
		UpdateContext: resourceTemplateUpdate,
		DeleteContext: resourceTemplateDelete,
		Schema:        templateSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	tflog.Info(ctx, "Create template from vm")
	activity, err := c.CreateTemplateFromVM(d.Get("fromvm").(int), d.Get("name").(string))

	if err != nil {
		return diag.FromErr(err)
	}

	for activity.Succeeded == nil {
		time.Sleep(time.Second)
		activity, err = c.GetInstanceActivities(activity.ID)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Info(ctx, fmt.Sprintf("Creating (%v) ... ", activity.GetPercentage))
	}
	if !*activity.Succeeded {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error while downloading ...",
			Detail:   "",
		})
		return diags
	}

	d.SetId(strconv.Itoa(activity.ResultData.Params.TemplateID))

	resourceTemplateRead(ctx, d, m)

	return diags
}

func resourceTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	TemplateID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Get template data from remote host")
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

	return diags
}

func resourceTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// TODO: update
	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	template_id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Delete template")
	err = c.DeleteTemplate(template_id)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
