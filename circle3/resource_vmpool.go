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

func resourceVMPool() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMPoolCreate,
		ReadContext:   resourceVMPoolRead,
		UpdateContext: resourceVMPoolUpdate,
		DeleteContext: resourceVMPoolDelete,
		Schema:        vmpoolSchema(),
	}
}

func flattenVM(vms *[]circleclient.VM) []interface{} {
	if vms != nil {
		vmsi := make([]interface{}, len(*vms))

		for i, vmitem := range *vms {
			vm := make(map[string]interface{})

			vm["id"] = strconv.Itoa(vmitem.ID)
			vm["name"] = vmitem.Name
			vm["pw"] = vmitem.Pw
			vm["status"] = vmitem.Status
			vm["node"] = vmitem.Node
			vm["ipv4"] = vmitem.Ipv4Addr
			vm["ipv6"] = vmitem.Ipv6Addr
			vm["disks"] = vmitem.Disks
			vm["vlans"] = vmitem.Vlans
			vm["cloud_init"] = vmitem.CloudInit
			vm["ci_user_data"] = vmitem.CiUserData
			vm["ci_meta_data"] = vmitem.CiMetaData
			vm["system"] = vmitem.System
			vm["has_agent"] = vmitem.HasAgent
			vm["num_cores"] = vmitem.NumCores
			vm["ram_size"] = vmitem.RamSize
			vm["max_ram_size"] = vmitem.MaxRamSize
			vm["arch"] = vmitem.Arch
			vm["priority"] = vmitem.Priority
			vm["sshportipv4"] = vmitem.SSHPortIpv4
			vm["hostipv4"] = vmitem.HostIpv4

			vmsi[i] = vm
		}
		return vmsi
	}
	return nil
}

func resourceVMPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics
	tflog.Info(ctx, "Create vmpool")

	from_template := d.Get("from_template").(int)
	name := d.Get("name").(string)
	users := d.Get("users").([]interface{})
	users_id := make([]int, len(users))
	for i, e := range users {
		users_id[i] = e.(int)
	}

	newvm, err := c.CreateVMfromTemplateforUsers(from_template, name, users_id)
	if err != nil {
		return diag.FromErr(err)
	}
	vmitems := flattenVM(&newvm)
	if err := d.Set("vms", vmitems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(time.Now().String())

	return diags
}

func resourceVMPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}

func resourceVMPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceVMPoolRead(ctx, d, m)
}

func resourceVMPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	if d.Get("vms") != nil {
		resource_vms := d.Get("vms").([]interface{})
		for _, e := range resource_vms {
			item := e.(map[string]interface{})
			id, err := strconv.Atoi(item["id"].(string))
			if err != nil {
				return diag.FromErr(err)
			}
			if err := c.DeleteVM(id); err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Problems with a VM (id %v)", id),
					Detail:   err.Error(),
				})
			}
		}
	}
	d.SetId("")

	return diags
}
