package provider

import (
	"context"
	"errors"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMUpdate,
		DeleteContext: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pw": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipv4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"access_method": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_menu": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"lease": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"cloud_init": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"ci_meta_data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ci_user_data": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system": {
				Type:     schema.TypeString,
				Required: true,
			},
			"has_agent": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"num_cores": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"ram_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_ram_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"arch": {
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"vlans": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional: true,
			},
			"disks": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional: true,
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	empty_req := []string{}
	vmrest := circleclient.VM{
		Owner:        d.Get("owner").(int),
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		Lease:        d.Get("lease").(int),
		CloudInit:    d.Get("cloud_init").(bool),
		CiMetaData:   d.Get("ci_meta_data").(string),
		CiUserData:   d.Get("ci_user_data").(string),
		System:       d.Get("system").(string),
		HasAgent:     d.Get("has_agent").(bool),
		NumCores:     d.Get("num_cores").(int),
		RamSize:      d.Get("ram_size").(int),
		MaxRamSize:   d.Get("max_ram_size").(int),
		Priority:     d.Get("priority").(int),
		AccessMethod: d.Get("access_method").(string),
		Arch:         d.Get("arch").(string),
		BootMenu:     d.Get("boot_menu").(bool),
		ReqTraits:    empty_req,
	}

	if d.Get("vlans") != nil {
		resource_vlans := d.Get("vlans").([]interface{})
		vlans := make([]int, len(resource_vlans))
		for _, e := range resource_vlans {
			vlans = append(vlans, e.(int))
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

	return diags
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}
	vm, err := c.GetVM(vmid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("pw", vm.Pw)
	d.Set("status", vm.Status)
	d.Set("node", strconv.Itoa(vm.Node))
	d.Set("ipv4", vm.Ipv4Addr)
	d.Set("ipv6", vm.Ipv6Addr)

	return diags
}

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	vmid, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChange("lease") {
		c.UpdateVMLease(vmid, d.Get("lease").(int))
	}

	if d.HasChange("state") {
		old, new := d.GetChange("state")
		c.UpdateVMState(vmid, old.(string), new.(string))
	}

	if d.HasChange("num_cores") || d.HasChange("ram_size") || d.HasChange("max_ram_size") || d.HasChange("priority") {
		if d.Get("state") == "STOPPED" {
			update := circleclient.VMResource{
				MaxRamSize: d.Get("max_ram_size").(int),
				RamSize:    d.Get("ram_size").(int),
				NumCores:   d.Get("num_cores").(int),
				Priority:   d.Get("priority").(int),
			}
			c.UpdateVMResource(vmid, update)
		} else {
			return diag.FromErr(errors.New("VM state is incorrect for change resources"))
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

	err = c.DeleteVM(vmid)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")

	return diags
}
