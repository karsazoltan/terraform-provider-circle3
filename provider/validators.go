package provider

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func ValidateSize(v interface{}, p cty.Path) diag.Diagnostics {
	error_msg := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Use this format: [positive integer][size unit], example: 10GB (units: GB, Gi, MB, Mi, KB, Ki)",
		Detail:   "",
	}
	value := v.(string)
	re := regexp.MustCompile(`(?P<size>\d*)\s*(?P<unit>\w*)`)
	matches := re.FindAllStringSubmatch(value, 1)
	var units = [...]string{"GB", "Gi", "MB", "Mi", "KB", "Ki"}
	var diags diag.Diagnostics

	size, err := strconv.Atoi(matches[0][1])
	if err != nil {
		diags = append(diags, error_msg)
		return diags
	}

	if size <= 0 {
		diags = append(diags, error_msg)
		return diags
	}

	for _, e := range units {
		if e == matches[0][2] {
			return diags
		}
	}
	return append(diags, error_msg)
}
