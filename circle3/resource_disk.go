package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceDisk() *schema.Resource {
	disk_schema := diskSchema()
	return &schema.Resource{
		CreateContext: resourceDiskCreate,
		ReadContext:   resourceDiskRead,
		UpdateContext: resourceDiskUpdate,
		DeleteContext: resourceDiskDelete,
		Schema:        disk_schema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceDiskCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if _, ok := d.GetOk("url"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceDDiskCreate(ctx, d, m)
		} else {
			return resourcePersistentDDiskCreate(ctx, d, m)
		}
	} else if _, ok := d.GetOk("size_format"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceCDiskCreate(ctx, d, m)
		} else {
			return resourcePersistentCDiskCreate(ctx, d, m)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One required field: url (download image from url) or size_format (empty disk)",
			Detail:   "",
		})
	}

	return diags
}

func resourceDiskRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if _, ok := d.GetOk("url"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceDDiskRead(ctx, d, m)
		} else {
			return resourcePersistentDDiskRead(ctx, d, m)
		}
	} else if _, ok := d.GetOk("size_format"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceCDiskRead(ctx, d, m)
		} else {
			return resourcePersistentCDiskRead(ctx, d, m)
		}
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One of size_format or url required",
			Detail:   "",
		})
	}

	return diags
}

func resourceDiskUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, ok := d.GetOk("url"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceDDiskUpdate(ctx, d, m)
		} else {
			return resourcePersistentDDiskUpdate(ctx, d, m)
		}
	} else if _, ok := d.GetOk("size_format"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceCDiskUpdate(ctx, d, m)
		} else {
			return resourcePersistentCDiskUpdate(ctx, d, m)
		}
	}
	return nil
}

func resourceDiskDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, ok := d.GetOk("url"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceDDiskDelete(ctx, d, m)
		} else {
			return resourcePersistentDDiskDelete(ctx, d, m)
		}
	} else if _, ok := d.GetOk("size_format"); ok {
		if _, ok = d.GetOk("vm"); ok {
			return resourceCDiskDelete(ctx, d, m)
		} else {
			return resourcePersistentCDiskDelete(ctx, d, m)
		}
	}
	return nil
}
