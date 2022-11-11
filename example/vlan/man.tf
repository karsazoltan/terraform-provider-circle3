terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}
provider "circle3" {
  address = "https://cloud3.fured.cloud.bme.hu"
  port    = 443
  // export CIRCLE3_TOKEN="secret-key"
}
variable "vlans" {
    type = map(object({
        user = string
    }))
}
data "circle3_vlan" "default_vlan" {
    for_each = var.vms
    name = 
}
resource "circle3_vm" "cluster" {
    for_each = var.vms
    name = "${each.key}"
    owner = each.value.user
    vlans = [data.circle3_vlan.labv[each.value.vlan]]
    from_template = data.circle3_template.basetemplate.id
}