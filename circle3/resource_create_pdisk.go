package provider

import (
	"context"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePersistentCDisk() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePersistentCDiskCreate,
		ReadContext:   resourcePersistentCDiskRead,
		UpdateContext: resourcePersistentCDiskUpdate,
		DeleteContext: resourcePersistentCDiskDelete,
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
		},
	}
}

func resourcePersistentCDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	cdisk := circleclient.CDisk{
		Size: d.Get("size_format").(string),
		Name: d.Get("name").(string),
	}

	disk, err := c.CreatePersistentCDisk(cdisk)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(disk.ID))

	resourcePersistentCDiskRead(ctx, d, m)

	return diags
}

func resourcePersistentCDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourcePersistentCDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePersistentCDiskRead(ctx, d, m)
}

func resourcePersistentCDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeletePersistentDisk(diskid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
