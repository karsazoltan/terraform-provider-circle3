package provider

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLeases() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceLeasesRead,
		Schema: map[string]*schema.Schema{
			"leases": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"suspend_interval_seconds": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"delete_interval_seconds": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceLeasesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}
	client := &http.Client{Timeout: 10 * time.Second, Transport: transCfg}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dashboard/acpi/lease/", "https://cloud3.fured.cloud.bme.hu"), nil)
	if err != nil {
		return diag.FromErr(err)
	}
	req.Header.Set("Authorization", "token 870d52e79fef266daebd1e6f781fe2c2422fde4a")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()
	leases := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&leases)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("leases", leases); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
