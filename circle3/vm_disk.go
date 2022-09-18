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

func resourceCDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	cdisk := circleclient.CDisk{
		Instance: d.Get("vm").(int),
		Size:     d.Get("size_format").(string),
		Name:     d.Get("name").(string),
	}
	tflog.Info(ctx, "Create disk for vm (download from url)")
	disk, err := c.CreateCDisk(cdisk)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(disk.ID))
	resourceCDiskRead(ctx, d, m)
	return diags
}

func resourceDDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	vmrest := circleclient.DDisk{
		Instance: d.Get("vm").(int),
		Url:      d.Get("url").(string),
		Name:     d.Get("name").(string),
	}
	tflog.Info(ctx, "Create disk for vm (empty disk)")
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

func resourceCDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Read cdisk")
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
	tflog.Info(ctx, "Delete cdisk")
	err = c.DeleteDisk(vmid, diskid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}

func resourceDDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	diskID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Read ddisk")
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

	tflog.Info(ctx, "Delete vm disk")
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
