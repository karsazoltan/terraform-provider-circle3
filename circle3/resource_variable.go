package provider

import (
	"context"
	"strconv"
	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVariableCreate,
		ReadContext:   resourceVariableRead,
		UpdateContext: resourceVariableUpdate,
		DeleteContext: resourceVariableDelete,
		Schema:        variableSchema(),
	}
}

func resourceVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	reqvar := circleclient.Variable{
		Key:   d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	respvar, err := c.CreateVariable(reqvar)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(respvar.ID))
	d.Set("url", respvar.URL)
	return diags
}

func resourceVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	varid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Read ddisk")
	respvar, err := c.GetVariablesByID(varid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", respvar.Key)
	d.Set("value", respvar.Value)
	d.Set("url", respvar.URL)

	return diags
}

func resourceVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceVariableRead(ctx, d, m)
}

func resourceVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	varid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "delete variable")
	c.DeleteVariable(varid)
	d.SetId("")

	return diags
}
