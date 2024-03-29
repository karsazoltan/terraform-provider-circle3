package provider

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMUpdate,
		DeleteContext: resourceVMDelete,
		Schema:        vmSchema(),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	if _, ok := d.GetOk("from_template"); ok {
		return resourceVMfromTemplateCreate(ctx, d, m)
	} else if _, ok := d.GetOk("from_rp"); ok {
		return resourceVMfromRP(ctx, d, m)
	}
	return resourceBaseVMCreate(ctx, d, m)
}

func resourceVMfromRP(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	tflog.Info(ctx, "Create vm from RP")
	if key, ok := d.GetOk("key"); ok {
		rpname := d.Get("from_rp").(string)
		newvm, err := c.CreateVMfromRP(rpname, key.(string))
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.Itoa(newvm.ID))
		return resourceVMRead(ctx, d, m)
	}
	var diags diag.Diagnostics
	return diags
}

func resourceBaseVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	tflog.Info(ctx, "Create base vm")
	empty_req := []string{}
	vmrest := circleclient.VM{
		Status:          d.Get("status").(string),
		Owner:           d.Get("owner").(int),
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Lease:           d.Get("lease").(int),
		CloudInit:       d.Get("cloud_init").(bool),
		CiMetaData:      d.Get("ci_meta_data").(string),
		CiUserData:      d.Get("ci_user_data").(string),
		CiNetworkConfig: d.Get("ci_network_config").(string),
		HookUrl:         d.Get("hookurl").(string),
		System:          d.Get("system").(string),
		HasAgent:        d.Get("has_agent").(bool),
		NumCores:        d.Get("num_cores").(int),
		NumCoresMax:     d.Get("num_cores_max").(int),
		RamSize:         d.Get("ram_size").(int),
		MaxRamSize:      d.Get("max_ram_size").(int),
		Priority:        d.Get("priority").(int),
		AccessMethod:    d.Get("access_method").(string),
		Arch:            d.Get("arch").(string),
		BootMenu:        d.Get("boot_menu").(bool),
		ReqTraits:       empty_req,
	}

	if d.Get("vlans") != nil {
		resource_vlans := d.Get("vlans").([]interface{})
		vlans := make([]int, len(resource_vlans))
		for i, e := range resource_vlans {
			vlans[i] = e.(int)
		}
		vmrest.Vlans = vlans
	}

	if d.Get("disks") != nil {
		resource_disks := d.Get("disks").([]interface{})
		disks := make([]int, len(resource_disks))
		for i, e := range resource_disks {
			disks[i] = e.(int)
		}
		vmrest.Disks = disks
	}

	newvm, err := c.CreateVM(vmrest)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(newvm.ID))

	return resourceVMRead(ctx, d, m)
}

func resourceVMfromTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)

	template_id := d.Get("from_template").(int)
	name := d.Get("name").(string)
	tflog.Info(ctx, "Create vm from template")
	newvm, err := c.CreateVMfromTemplate(template_id, name)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(strconv.Itoa(newvm.ID))
	return resourceVMRead(ctx, d, m)
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Get VM from remote host")
	vm, err := c.GetVM(vmid)
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

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, "Get VM from remote host")
	vm_remote, err := c.GetVM(vmid)
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("lease") {
		tflog.Info(ctx, "Update lease")
		c.UpdateVMLease(vmid, d.Get("lease").(int))
	}

	if d.HasChange("status") {
		old, new := d.GetChange("status")
		if old.(string) != vm_remote.Status {
			tflog.Warn(ctx, "Remote vm status and local state is inconsistent!")
			c.UpdateVMState(vmid, vm_remote.Status, new.(string))
		} else {
			tflog.Info(ctx, fmt.Sprintf("Update vm status: %s -> %s", old.(string), new.(string)))
			c.UpdateVMState(vmid, old.(string), new.(string))
		}
	}

	if d.Get("status").(string) != "RUNNING" {
		if d.HasChange("max_ram_size") || d.HasChange("priority") ||
			d.HasChange("num_cores_max") || d.HasChange("num_cores") || d.HasChange("ram_size") {
			tflog.Info(ctx, "Update vm static resources")
			update := circleclient.VMResource{
				MaxRamSize:  d.Get("max_ram_size").(int),
				RamSize:     d.Get("ram_size").(int),
				NumCores:    d.Get("num_cores").(int),
				Priority:    d.Get("priority").(int),
				NumCoresMax: d.Get("num_cores_max").(int),
			}
			c.UpdateVMResource(vmid, update)
		}
	} else {
		if d.HasChange("num_cores") {
			tflog.Info(ctx, "Update vm hotplug num_cores")
			c.UpdateVCPUHotplug(vmid, d.Get("num_cores").(int))
		}
		if d.HasChange("ram_size") {
			tflog.Info(ctx, "Update vm hotplug ram_size")
			c.UpdateMemHotplug(vmid, d.Get("ram_size").(int))
		}
		if d.HasChange("max_ram_size") || d.HasChange("priority") || d.HasChange("num_cores_max") {
			return diag.FromErr(errors.New("VM state is incorrect for change resources"))
		}
	}

	if d.HasChange("disks") {
		old, new := d.GetChange("disks")
		old_disks := old.([]interface{})
		new_disks := new.([]interface{})

		olds_int := make([]int, 0)
		news_int := make([]int, 0)
		for _, n := range new_disks {
			news_int = append(news_int, n.(int))
		}
		for _, n := range old_disks {
			olds_int = append(olds_int, n.(int))
		}

		for _, n := range news_int {
			// new disks
			if !contains(olds_int, n) {
				tflog.Info(ctx, fmt.Sprintf("Add persistent disk (%v) to vm", n))
				c.AddNewPersistentDiskToVM(vmid, n)
			}
		}
		for _, n := range olds_int {
			// deleted disks
			if !contains(news_int, n) {
				tflog.Info(ctx, fmt.Sprintf("Delete persistent disk (%v) from vm", n))
				c.DeleteDisk(vmid, n)
			}
		}
	}

	return resourceVMRead(ctx, d, m)
}

func resourceVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	tflog.Info(ctx, fmt.Sprintf("Delete vm (%v)", vmid))
	err = c.DeleteVM(vmid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
