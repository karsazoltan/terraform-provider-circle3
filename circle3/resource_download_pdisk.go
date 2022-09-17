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

func resourcePersistentDDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePersistentDDiskCreate,
		ReadContext:   resourcePersistentDDiskRead,
		UpdateContext: resourcePersistentDDiskUpdate,
		DeleteContext: resourcePersistentDDiskDelete,
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
		},
	}
}

func resourcePersistentDDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmrest := circleclient.DDisk{
		Url:  d.Get("url").(string),
		Name: d.Get("name").(string),
	}
	tflog.Info(ctx, "Create persistent disk (download from url)")
	activity, err := c.CreatePersistentDDisk(vmrest)

	if err != nil {
		return diag.FromErr(err)
	}

	for !activity.Succeeded {
		time.Sleep(time.Second)
		activity, err = c.GetStorageActivity(activity.ID)
		if err != nil {
			return diag.FromErr(err)
		}
		tflog.Info(ctx, "Downloading ... ")
	}

	d.SetId(strconv.Itoa(activity.Disk))

	resourcePersistentDDiskRead(ctx, d, m)

	return diags
}

func resourcePersistentDDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	tflog.Info(ctx, "Read remote host: persistent ddisk")
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

func resourcePersistentDDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePersistentDDiskRead(ctx, d, m)
}

func resourcePersistentDDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Delete persistent disk")
	err = c.DeletePersistentDisk(diskid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
