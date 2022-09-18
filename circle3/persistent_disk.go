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

func resourcePersistentCDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	cdisk := circleclient.CDisk{
		Size: d.Get("size_format").(string),
		Name: d.Get("name").(string),
	}
	tflog.Info(ctx, "Create persistent disk (empty disk)")
	disk, err := c.CreatePersistentCDisk(cdisk)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(disk.ID))
	resourcePersistentCDiskRead(ctx, d, m)
	return diags
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
