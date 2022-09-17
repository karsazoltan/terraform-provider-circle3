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

func ValidatePositiveNumber(v interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	number := v.(int)
	if number <= 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value must be a positive integer",
			Detail:   "",
		})
	}

	return diags
}

func CheckPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func ValidateRamNumber(v interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	ValidatePositiveNumber(v, p)
	number := v.(int)

	if CheckPowerOfTwo(number) != 0 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Value must be power of 2",
			Detail:   "",
		})
		return diags
	}

	return diags
}
