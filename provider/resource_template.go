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
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func resourceTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	activity, err := c.CreateTemplateFromVM(d.Get("fromvm").(int), d.Get("name").(string))

	if err != nil {
		return diag.FromErr(err)
	}

	for !activity.Succeeded {
		time.Sleep(time.Second)
		activity, err = c.GetInstanceActivities(activity.ID)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Info(ctx, fmt.Sprintf("Downloading (%v) ... ", activity.GetPercentage))
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
	return resourceTemplateRead(ctx, d, m)
}

func resourceTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	vmid := d.Get("vm").(int)

	err = c.DeleteDisk(vmid, diskid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
