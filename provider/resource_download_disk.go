package provider

import (
	"context"
	"log"
	"strconv"
	"time"

	circleclient "terraform-provider-circle3/client"

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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"checksum": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vm": &schema.Schema{
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

	for activity.Succeeded != true {
		time.Sleep(time.Second)
		activity, err = c.GetActivity(activity.ID)
		if err != nil {
			return diag.FromErr(err)
		}
		log.Printf("Downloading (%v%) ... ", activity.GetPercentage)
	}

	d.SetId(strconv.Itoa(activity.ResultData.Params.DiskID))
	d.Set("size", activity.ResultData.Params.DiskSize)
	d.Set("checksum", activity.ResultData.Params.Checksum)

	return diags
}

func resourceDDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceDDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceDDiskRead(ctx, d, m)
}

func resourceDDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	return diags
}
