package provider

import (
	"context"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCDiskCreate,
		ReadContext:   resourceCDiskRead,
		UpdateContext: resourceCDiskUpdate,
		DeleteContext: resourceCDiskDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size_format": {
				Type:             schema.TypeString,
				Required:         true,
				ValidateDiagFunc: ValidateSize,
			},
			"vm": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceCDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	cdisk := circleclient.CDisk{
		Instance: d.Get("vm").(int),
		Size:     d.Get("size_format").(string),
		Name:     d.Get("name").(string),
	}

	disk, err := c.CreateCDisk(cdisk)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(disk.ID))

	resourceCDiskRead(ctx, d, m)

	return diags
}

func resourceCDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceCDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceCDiskRead(ctx, d, m)
}

func resourceCDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
