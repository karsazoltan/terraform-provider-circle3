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
data "circle3_template" "meres" {
  name = "meres-sablon"
}
variable "users" {
  type = map(list(string))
  default = {
    "vm" = [ "Alice", "Joe" ]
    "eth0" = [ "Bob" ]
  }
}
data "circle3_vlan" "meres_vlan" {
  for_each = var.users
  name = "${each.key}"
}
locals  {
  vms = flatten([ 
    for k, v in var.users : [ for u in v: { user = u, vlan = k}  ]
  ])
} 
resource "circle3_vm" "meres_vms" {
  for_each = {for vm in local.vms: 
    "${vm.user}-VM" => { vlan = vm.vlan, owner = vm.user 
  }}
  name = "${each.key}"
  owner = each.value.owner
  vlans = [data.circle3_vlan.meres_vlan[each.value.vlan]]
  from_template = data.circle3_template.meres.id
}