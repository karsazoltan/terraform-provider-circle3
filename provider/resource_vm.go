package provider

import (
	"context"
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"node": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"access_method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"boot_menu": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"lease": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"cloud_init": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"ci_meta_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ci_user_data": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"system": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"has_agent": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"num_cores": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"ram_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"max_ram_size": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"arch": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	empty_disk := []int{}
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
		Disks:        empty_disk,
		ReqTraits:    empty_req,
	}

	newvm, err := c.CreateVM(vmrest)

	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(newvm.ID))
	d.Set("pw", newvm.Pw)
	d.Set("status", newvm.Status)
	d.Set("node", strconv.Itoa(newvm.Node))

	return diags
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
