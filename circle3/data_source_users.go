package provider

import (
	"context"
	"strconv"

	circleclient "terraform-provider-circle3/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"firs_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_superuser": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_staff": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"groups": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*circleclient.Client)
	var diags diag.Diagnostics

	if _, ok := d.GetOk("username"); ok {
		tflog.Info(ctx, "Get user by username")
		user, err := c.GetUserByName(d.Get("username").(string))
		if err != nil {
			return diag.FromErr(err)
		}

		d.SetId(strconv.Itoa(user.ID))
		d.Set("email", user.Email)
		d.Set("is_staff", user.IsStaff)
		d.Set("is_superuser", user.IsSuperuser)
		d.Set("first_name", user.FirstName)
		d.Set("last_name", user.LastName)
		d.Set("groups", user.Groups)
	} else if _, ok := d.GetOk("id"); ok {
		tflog.Info(ctx, "Get user by id")
		id, err := strconv.Atoi(d.Id())
		if err != nil {
			return diag.FromErr(err)
		}
		user, err := c.GetUserByID(id)
		if err != nil {
			return diag.FromErr(err)
		}
		d.SetId(strconv.Itoa(user.ID))
		d.Set("username", user.Username)
		d.Set("email", user.Email)
		d.Set("is_staff", user.IsStaff)
		d.Set("is_superuser", user.IsSuperuser)
		d.Set("first_name", user.FirstName)
		d.Set("last_name", user.LastName)
		d.Set("groups", user.Groups)
	} else {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "One required field: username or id",
			Detail:   "",
		})
	}

	return diags
}
