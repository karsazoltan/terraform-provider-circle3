terraform {
  required_providers {
    circle3 = {
      version = "0.1"
      source  = "bmeik/tf/circle3"
    }
  }
}

provider "circle3" {
  address = "https://meres.fured.cloud.bme.hu"
  port    = 443
  //token   = "secret" -> export CIRCLE3_TOKEN="..."
}
variable "list" {
  type = list(string)
}
data "circle3_template" "basetemplate" {
  name = "meres-sablon"
}
resource "circle3_vm" "for_each_users" {
  for_each = toset(var.list)
  name = "vm pool ${each.key}"
  from_template = data.circle3_template.basetemplate.id
  owner = 1
}