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
  type = list(number)
}
data "circle3_template" "basetemplate" {
  name = "meres-sablon"
}
resource "circle3_vmpool" "pool_users" {
  name = "vm pool"
  from_template = data.circle3_template.basetemplate.id
  users = var.list
}