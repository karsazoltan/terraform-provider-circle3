---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "circle3_vm Resource - terraform-provider-circle3"
subcategory: ""
description: |-
  
---

# circle3_vm (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String)

### Optional

- `access_method` (String)
- `arch` (String)
- `boot_menu` (Boolean)
- `ci_meta_data` (String)
- `ci_user_data` (String)
- `cloud_init` (Boolean)
- `description` (String)
- `disks` (List of Number)
- `from_template` (Number)
- `has_agent` (Boolean)
- `lease` (Number)
- `max_ram_size` (Number)
- `num_cores` (Number)
- `owner` (Number)
- `priority` (Number)
- `ram_size` (Number)
- `status` (String)
- `system` (String)
- `vlans` (List of Number)

### Read-Only

- `id` (String) The ID of this resource.
- `ipv4` (String)
- `ipv6` (String)
- `node` (Number)
- `pw` (String)


