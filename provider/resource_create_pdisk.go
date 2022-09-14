package provider

import (
	"context"
	"regexp"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/go-cty/cty"
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
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"filename": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"datastore": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bus": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"base": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dev_num": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destroyed": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ci_disk": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_ready": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size_format": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					error_msg := diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Use this format: [positive integer][size unit], example: 10GB (units: GB, Gi, MB, Mi, KB, Ki)",
						Detail:   "",
					}
					value := v.(string)
					re := regexp.MustCompile(`(?P<size>\d*)\s*(?P<unit>\w*)`)
					matches := re.FindAllStringSubmatch(value, 1)
					var units = [...]string{"GB", "Gi", "MB", "Mi", "KB", "Ki"}
					var diags diag.Diagnostics

					size, err := strconv.Atoi(matches[0][1])
					if err != nil {
						diags = append(diags, error_msg)
						return diags
					}

					if size <= 0 {
						diags = append(diags, error_msg)
						return diags
					}

					for _, e := range units {
						if e == matches[0][2] {
							return diags
						}
					}
					return append(diags, error_msg)
				},
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
	vmid := d.Get("vm").(int)

	err = c.DeleteDisk(vmid, diskid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
