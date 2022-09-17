package provider

import (
	"context"
	"fmt"
	"strconv"
	"time"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDDiskCreate,
		ReadContext:   resourceDDiskRead,
		UpdateContext: resourceDDiskUpdate,
		DeleteContext: resourceDDiskDelete,
		Schema: map[string]*schema.Schema{
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
			"checksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"vm": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceDDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmrest := circleclient.DDisk{
		Instance: d.Get("vm").(int),
		Url:      d.Get("url").(string),
		Name:     d.Get("name").(string),
	}

	activity, err := c.CreateDDisk(vmrest)

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

	d.SetId(strconv.Itoa(activity.ResultData.Params.DiskID))
	d.Set("checksum", activity.ResultData.Params.Checksum)

	resourceDDiskRead(ctx, d, m)

	return diags
}

func resourceDDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	disk, err := c.GetDisk(diskID)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("size", disk.Size)
	d.Set("filename", disk.Filename)
	d.Set("datastore", disk.Datastore)
	d.Set("type", disk.Type)
	d.Set("bus", disk.Bus)
	d.Set("base", disk.Base)
	d.Set("dev_num", disk.DevNum)
	d.Set("destroyed", disk.Destroyed)
	d.Set("ci_disk", disk.CiDisk)
	d.Set("is_ready", disk.IsReady)

	return diags
}

func resourceDDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDDiskRead(ctx, d, m)
}

func resourceDDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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