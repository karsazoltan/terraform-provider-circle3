package provider

import (
	"context"
	"fmt"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLBVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLBVMCreate,
		ReadContext:   resourceLBVMRead,
		UpdateContext: resourceLBVMUpdate,
		DeleteContext: resourceLBVMDelete,
		Schema:        vmLoadBalancingSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceLBVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceLBVMfromTemplateCreate(ctx, d, m)
}

func resourceLBVMfromTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	template_name := d.Get("from_template").(string)
	name := d.Get("name").(string)
	username := d.Get("username").(string)
	balancermod := d.Get("balancer_method").(string)
	tflog.Info(ctx, "Create vm from template")
	newvm, err := c.CreateLBVMfromTemplate(template_name, name, username, balancermod)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(newvm.ID))
	d.Set("datacenter", newvm.DataCenter)
	return resourceLBVMRead(ctx, d, m)
}

func resourceLBVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceLBVMRead(ctx, d, m)
}

func resourceLBVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	datacenter := d.Get("datacenter").(string)
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Get VM from remote host")
	vm, err := c.GetLBVM(vmid, datacenter)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("pw", vm.Pw)
	d.Set("status", vm.Status)
	d.Set("node", strconv.Itoa(vm.Node))
	d.Set("ipv4", vm.Ipv4Addr)
	d.Set("ipv6", vm.Ipv6Addr)
	d.Set("disks", vm.Disks)
	d.Set("vlans", vm.Vlans)
	d.Set("cloud_init", vm.CloudInit)
	d.Set("ci_user_data", vm.CiUserData)
	d.Set("ci_meta_data", vm.CiMetaData)
	d.Set("ci_network_config", vm.CiNetworkConfig)
	d.Set("system", vm.System)
	d.Set("has_agent", vm.HasAgent)
	d.Set("num_cores", vm.NumCores)
	d.Set("num_cores_max", vm.NumCoresMax)
	d.Set("ram_size", vm.RamSize)
	d.Set("max_ram_size", vm.MaxRamSize)
	d.Set("arch", vm.Arch)
	d.Set("priority", vm.Priority)
	d.Set("sshportipv4", vm.SSHPortIpv4)
	d.Set("hostipv4", vm.HostIpv4)

	return diags
}

func resourceLBVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	datacenter := d.Get("datacenter").(string)
	tflog.Info(ctx, fmt.Sprintf("Delete vm (%v)", vmid))
	err = c.DeleteLBVM(vmid, datacenter)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
