package provider

import (
	"context"
	"fmt"
	"terraform-provider-circle3/client"
	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePort() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePortCreate,
		ReadContext:   resourcePortRead,
		UpdateContext: resourcePortUpdate,
		DeleteContext: resourcePortDelete,
		Schema:        portSchema(),
	}
}

func resourcePortCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	interf := client.PortsReq{
		Vlan:     d.Get("vlan").(int),
		Instance: d.Get("instance").(int),
	}
	port := client.OpenPort{
		DestinationPort: d.Get("port").(int),
		Type:            d.Get("type").(string),
	}

	portresponse, err := c.CreatePort(interf, port)
	if err != nil {
		diag.FromErr(err)
	}
	d.SetId(fmt.Sprintf("%v/%v/%v", interf.Instance, interf.Vlan, portresponse.DestinationPort))
	d.Set("source_port", portresponse.SourcePort)
	d.Set("forwarding", portresponse.Forwarding)
	return diags
}

func resourcePortRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourcePortUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourcePortRead(ctx, d, m)
}

func resourcePortDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")

	return diags
}
